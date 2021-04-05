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
	return fmt.Sprintf("[%s] (%d, %d) %s", TokenNames[t.TokenType], t.pos, t.length, strings.TrimSpace(t.Value))
}

type TokenType int

const (
	TokenError TokenType = iota
	TokenEOF
	TokenNewline

	TokenComment

	TokenIdentifier
	TokenLiteralString
	TokenLiteralInt
)

var TokenNames = map[TokenType]string{
	TokenError:   "Error",
	TokenEOF:     "EOF",
	TokenNewline: "Newline",

	TokenComment: "Comment",

	TokenIdentifier:    "Identifier",
	TokenLiteralString: "String",
	TokenLiteralInt:    "Int",
}

const (
	eof          rune = 0
	commentStart rune = '#'
	newline      rune = '\n'
	quote        rune = '"'
)
