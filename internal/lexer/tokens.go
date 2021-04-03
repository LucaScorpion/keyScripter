package lexer

type token struct {
	tokenType tokenType
	value     string
}

type tokenType int

const (
	tokenError tokenType = iota
	tokenEOF

	tokenComment

	tokenIdentifier
	tokenValue
)

const (
	eof          rune   = 0
	commentStart string = "#"
	newline      rune   = '\n'
)
