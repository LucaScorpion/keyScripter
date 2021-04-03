package lexer

import (
	"fmt"
	"strings"
	"unicode"
)

type lexer struct {
	input  string
	Tokens chan *Token
	state  lexFn

	start   int
	current int
}

type lexFn func(*lexer) lexFn

func NewLexer(input string) *lexer {
	return &lexer{
		input:  strings.ReplaceAll(input, "\r\n", string(newline)),
		Tokens: make(chan *Token),
		state:  lexBegin,
	}
}

// Start the lexer, emitting tokens to the Tokens channel.
func (l *lexer) Run() {
	for state := l.state; state != nil; {
		state = state(l)
	}
	close(l.Tokens)
}

// Emit the current token in the Tokens channel.
func (l *lexer) emit(tokenType TokenType) {
	if l.current == l.start && tokenType != TokenEOF {
		panic(fmt.Errorf("cannot emit token when current == start, at pos %d", l.start))
	}

	l.Tokens <- &Token{
		TokenType: tokenType,
		Value:     l.input[l.start:l.current],
		pos:       l.start,
		length:    l.current - l.start,
	}
	l.start = l.current
}

// Discard the current token by setting the start position to current.
func (l *lexer) discard() {
	l.start = l.current
}

// Emit an error token.
func (l *lexer) errorf(err string, args ...interface{}) lexFn {
	l.Tokens <- &Token{
		TokenType: TokenError,
		Value:     fmt.Sprintf(err, args...),
		pos:       l.current,
		length:    0,
	}
	return nil
}

func (l *lexer) peekRune() rune {
	if l.current >= len(l.input) {
		return eof
	}
	return rune(l.input[l.current])
}

func (l *lexer) readRune() rune {
	nextRune := l.peekRune()
	l.current++
	return nextRune
}

func (l *lexer) unreadRune() {
	if l.current <= l.start {
		panic("cannot unread rune when current <= start")
	}
	l.current--
}

func (l *lexer) currentInput() string {
	return l.input[l.current:]
}

func (l *lexer) readWhile(whileFn func(rune) bool) {
	for {
		nextRune := l.readRune()
		if nextRune == eof || !whileFn(nextRune) {
			l.unreadRune()
			break
		}
	}
}

// Read spaces.
func (l *lexer) readSpace(includeNewline bool) {
	l.readWhile(func(r rune) bool {
		return unicode.IsSpace(r) && (includeNewline || r != newline)
	})
}

// Read until the end of line.
func (l *lexer) readLine() {
	l.readWhile(func(r rune) bool {
		return r != newline
	})
}

// Read alphanumeric runes.
func (l *lexer) readAlphaNum() {
	l.readWhile(func(r rune) bool {
		return unicode.IsLetter(r) || unicode.IsNumber(r)
	})
}

// Expect a space or EOF.
// If the next rune is anything else, errorf is used to emit an error.
func (l *lexer) expectSpace() bool {
	nextRune := l.peekRune()
	if nextRune != eof && !unicode.IsSpace(nextRune) {
		l.errorf("unexpected character: %s", string(nextRune))
		return false
	}
	return true
}

func (l *lexer) readStringLiteral() {
	// Opening quote
	l.readRune()

	var nextRune rune
	escaped := false

	for {
		nextRune = l.readRune()

		if escaped {
			escaped = false
		} else if nextRune == '\\' {
			escaped = true
		} else if nextRune == quote {
			break
		}
	}
}
