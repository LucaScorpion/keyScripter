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
	return t.TokenType > beginValueTokens && t.TokenType < endValueTokens
}

func (t *Token) Name() string {
	return TokenName(t.TokenType)
}

func TokenName(t TokenType) string {
	return tokenNames[t]
}

type TokenType int

const (
	TokenEOF TokenType = iota
	TokenNewline

	TokenComment

	beginValueTokens
	TokenIdentifier
	TokenLiteralString
	TokenLiteralInt
	TokenLiteralHex
	endValueTokens

	TokenAssign
	TokenParenOpen
	TokenParenClose
	TokenBlockStart
	TokenBlockEnd
)

var tokenNames = map[TokenType]string{
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
	TokenBlockStart: "{",
	TokenBlockEnd:   "}",
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
	blockStart   rune = '{'
	blockEnd     rune = '}'
)
