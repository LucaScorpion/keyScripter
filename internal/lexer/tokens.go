package lexer

import "fmt"

type Token struct {
	TokenType TokenType
	Value     string
	pos       int
	length    int
}

func (t Token) String() string {
	return fmt.Sprintf("[%s] (%d, %d) %s", TokenNames[t.TokenType], t.pos, t.length, t.Value)
}

type TokenType int

const (
	TokenError TokenType = iota
	TokenEOF

	TokenComment

	TokenFuncName
	TokenLiteralString
	TokenLiteralInt
)

var TokenNames = map[TokenType]string{
	TokenError: "Error",
	TokenEOF:   "EOF",

	TokenComment: "Comment",

	TokenFuncName:      "Function",
	TokenLiteralString: "String",
	TokenLiteralInt:    "Int",
}

const (
	eof          rune   = 0
	commentStart string = "#"
	newline      rune   = '\n'
	quote        rune   = '"'
)
