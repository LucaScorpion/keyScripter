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
	return tokenNames[t.TokenType]
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
	tokenValueEnd

	TokenAssign
)

var tokenNames = map[TokenType]string{
	TokenError:   "Error",
	TokenEOF:     "EOF",
	TokenNewline: "Newline",

	TokenComment: "Comment",

	TokenIdentifier:    "Identifier",
	TokenLiteralString: "String",
	TokenLiteralInt:    "Int",

	TokenAssign: "Assign",
}

const (
	eof          rune = 0
	commentStart rune = '#'
	newline      rune = '\n'
	quote        rune = '"'
	equals       rune = '='
)
