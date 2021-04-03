package lexer

import "fmt"

type token struct {
	tokenType tokenType
	value     string
	pos       int
	length    int
}

func (t token) String() string {
	return fmt.Sprintf("[%s] (%d, %d) %s", tokenNames[t.tokenType], t.pos, t.length, t.value)
}

type tokenType int

const (
	tokenError tokenType = iota
	tokenEOF

	tokenComment

	tokenIdentifier
	tokenLiteralString
	tokenLiteralInt
)

var tokenNames = map[tokenType]string{
	tokenError: "Error",
	tokenEOF:   "EOF",

	tokenComment: "Comment",

	tokenIdentifier:    "Identifier",
	tokenLiteralString: "String",
	tokenLiteralInt:    "Int",
}

const (
	eof          rune   = 0
	commentStart string = "#"
	newline      rune   = '\n'
	quote        rune   = '"'
)
