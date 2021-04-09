package lexer

import (
	"fmt"
	"unicode"
)

var oneRuneTokens = map[rune]TokenType{
	newline:    TokenNewline,
	equals:     TokenAssign,
	parenOpen:  TokenParenOpen,
	parenClose: TokenParenClose,
	blockStart: TokenBlockStart,
	blockEnd:   TokenBlockEnd,
}

func lexBegin(l *lexer) (lexFn, error) {
	// Discard leading whitespaces.
	l.readSpace()
	l.discard()

	nextRune := l.peekRune()
	oneRuneToken, isOneRuneToken := oneRuneTokens[nextRune]

	switch {
	case nextRune == eof:
		l.emit(TokenEOF)
		return nil, nil
	case isOneRuneToken:
		l.readRune()
		l.emit(oneRuneToken)
		return lexBegin, nil
	case nextRune == commentStart:
		l.readLine()
		l.emit(TokenComment)
		return lexBegin, nil
	case unicode.IsLetter(nextRune) || nextRune == '_':
		l.readAlphaNum()
		l.emit(TokenIdentifier)
		return lexBegin, nil
	case unicode.IsNumber(nextRune):
		return lexNumberLiteral, nil
	case nextRune == quote:
		return lexStringLiteral, nil
	default:
		return nil, fmt.Errorf("unexpected character '%s' at position %d", string(nextRune), l.current)
	}
}

func lexStringLiteral(l *lexer) (lexFn, error) {
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

	l.emit(TokenLiteralString)
	return lexBegin, nil
}

func lexNumberLiteral(l *lexer) (lexFn, error) {
	// Leading number
	l.readRune()

	if l.peekRune() == hex {
		l.readRune()
		l.readWhile(unicode.IsNumber)
		l.emit(TokenLiteralHex)
	} else {
		l.readWhile(unicode.IsNumber)
		l.emit(TokenLiteralInt)
	}

	return lexBegin, nil
}
