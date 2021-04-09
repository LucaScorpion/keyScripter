package lexer

import (
	"fmt"
	"strings"
)

type Token struct {
	TokenType TokenType
	Value     string
	pos       int
	length    int
}

func (t Token) String() string {
	return fmt.Sprintf("[%s] (%d, %d) %s", t.Name(), t.pos, t.length, strings.TrimSpace(t.Value))
}

func (t *Token) IsValueToken() bool {
	return t.TokenType > tokenValueStart && t.TokenType < tokenValueEnd
}

func (t *Token) Name() string {
	return TokenName(t.TokenType)
}

func TokenName(t TokenType) string {
	return tokenNames[t]
}

type TokenType int

const (
	TokenError TokenType = iota
	TokenEOF
	TokenNewline

	TokenComment

	tokenValueStart
	TokenIdentifier
	TokenLiteralString
	TokenLiteralInt
	TokenLiteralHex
	tokenValueEnd

	TokenAssign
	TokenParenOpen
	TokenParenClose
	TokenBraceOpen
	TokenBraceClose
)

var tokenNames = map[TokenType]string{
	TokenError:   "Error",
	TokenEOF:     "EOF",
	TokenNewline: "Newline",

	TokenComment: "Comment",

	TokenIdentifier:    "Identifier",
	TokenLiteralString: "String",
	TokenLiteralInt:    "Int",
	TokenLiteralHex:    "Hex",

	TokenAssign:     "Assign",
	TokenParenOpen:  "(",
	TokenParenClose: ")",
	TokenBraceOpen:  "{",
	TokenBraceClose: "}",
}

const (
	eof          rune = 0
	commentStart rune = '#'
	newline      rune = '\n'
	quote        rune = '"'
	equals       rune = '='
	hex          rune = 'x'
	parenOpen    rune = '('
	parenClose   rune = ')'
	braceOpen    rune = '{'
	braceClose   rune = '}'
)
