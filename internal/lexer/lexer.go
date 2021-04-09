package lexer

import (
	"fmt"
	"strings"
	"unicode"
)

type lexFn func(*lexer) (lexFn, error)

type lexer struct {
	input  string
	tokens []*Token
	state  lexFn

	start   int
	current int
}

func NewLexer(input string) *lexer {
	return &lexer{
		input: strings.ReplaceAll(input, "\r", ""),
		state: lexBegin,
	}
}

func (l *lexer) Tokens() []*Token {
	return l.tokens
}

// Run the lexer.
func (l *lexer) Run() error {
	for l.state != nil {
		newState, err := l.state(l)
		if err != nil {
			return err
		}
		l.state = newState
	}
	return nil
}

// Store the current token in the tokens slice.
func (l *lexer) emit(tokenType TokenType) {
	if l.current == l.start && tokenType != TokenEOF {
		panic(fmt.Errorf("cannot emit token when current == start, at pos %d", l.start))
	}

	l.tokens = append(l.tokens, &Token{
		TokenType: tokenType,
		Value:     l.currentValue(),
		pos:       l.start,
		length:    l.current - l.start,
	})
	l.start = l.current
}

func (l *lexer) currentValue() string {
	return l.input[l.start:l.current]
}

// Discard the current token by setting the start position to current.
func (l *lexer) discard() {
	l.start = l.current
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

func (l *lexer) readWhile(whileFn func(rune) bool) {
	for {
		nextRune := l.readRune()
		if nextRune == eof || !whileFn(nextRune) {
			l.unreadRune()
			break
		}
	}
}

// Read spaces, excluding newlines.
func (l *lexer) readSpace() {
	l.readWhile(func(r rune) bool {
		return unicode.IsSpace(r) && r != newline
	})
}

// Read until the end of line.
func (l *lexer) readLine() {
	l.readWhile(func(r rune) bool {
		return r != newline
	})
}

// Read alphanumeric runes (a-z, 0-9, _).
func (l *lexer) readAlphaNum() {
	l.readWhile(func(r rune) bool {
		return unicode.IsLetter(r) || unicode.IsNumber(r) || r == '_'
	})
}
