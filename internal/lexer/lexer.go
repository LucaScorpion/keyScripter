package lexer

import (
	"fmt"
	"strings"
	"unicode"
)

type lexer struct {
	input  string
	Tokens chan token
	state  lexFn

	start   int
	current int
}

type lexFn func(*lexer) lexFn

func NewLexer(input string) *lexer {
	return &lexer{
		input:  strings.ReplaceAll(input, "\r\n", string(newline)),
		Tokens: make(chan token),
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
func (l *lexer) emit(tokenType tokenType) {
	l.Tokens <- token{
		tokenType: tokenType,
		value:     l.input[l.start:l.current],
	}
	l.start = l.current
}

// Discard the current token by setting the start position to current.
func (l *lexer) discard() {
	l.start = l.current
}

// Emit an error token.
func (l *lexer) errorf(err string, args ...interface{}) lexFn {
	l.Tokens <- token{
		tokenType: tokenError,
		value:     fmt.Sprintf(err, args...),
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

func (l *lexer) readWhile(whileFn func(rune) bool) string {
	val := ""
	for {
		nextRune := l.readRune()
		if nextRune == eof || !whileFn(nextRune) {
			l.unreadRune()
			break
		}
		val += string(nextRune)
	}
	return val
}

// Read spaces.
func (l *lexer) readSpace(includeNewline bool) string {
	return l.readWhile(func(r rune) bool {
		return unicode.IsSpace(r) && (includeNewline || r != newline)
	})
}

// Read until the end of line.
func (l *lexer) readLine() string {
	return l.readWhile(func(r rune) bool {
		return r != newline
	})
}

// Read alphanumeric runes.
func (l *lexer) readAlphaNum() string {
	return l.readWhile(func(r rune) bool {
		return unicode.IsLetter(r) || unicode.IsNumber(r)
	})
}
