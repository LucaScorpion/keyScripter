package lexer

import (
	"unicode"
)

var oneRuneTokens = map[rune]TokenType{
	newline:    TokenNewline,
	equals:     TokenAssign,
	parenOpen:  TokenParenOpen,
	parenClose: TokenParenClose,
	braceOpen:  TokenBraceOpen,
	braceClose: TokenBraceClose,
}

func lexBegin(l *lexer) lexFn {
	// Discard leading whitespaces.
	l.readSpace()
	l.discard()

	nextRune := l.peekRune()
	oneRuneToken, isOneRuneToken := oneRuneTokens[nextRune]

	switch {
	case nextRune == eof:
		l.emit(TokenEOF)
		return nil
	case isOneRuneToken:
		l.readRune()
		l.emit(oneRuneToken)
		return lexBegin
	case nextRune == commentStart:
		l.readLine()
		l.emit(TokenComment)
		return lexBegin
	case unicode.IsLetter(nextRune) || nextRune == '_':
		l.readAlphaNum()
		l.emit(TokenIdentifier)
		return lexBegin
	case unicode.IsNumber(nextRune):
		return lexNumberLiteral
	case nextRune == quote:
		return lexStringLiteral
	default:
		return l.errorf("unexpected character: '%s'", string(nextRune))
	}
}

func lexStringLiteral(l *lexer) lexFn {
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
	return lexBegin
}

func lexNumberLiteral(l *lexer) lexFn {
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

	return lexBegin
}
