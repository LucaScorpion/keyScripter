package lexer

import (
	"unicode"
)

func lexBegin(l *lexer) lexFn {
	// Discard leading whitespaces.
	l.readSpace()
	l.discard()

	nextRune := l.peekRune()
	switch {
	case nextRune == eof:
		l.emit(TokenEOF)
		return nil
	case nextRune == newline:
		l.readRune()
		l.emit(TokenNewline)
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
		l.readWhile(unicode.IsNumber)
		l.emit(TokenLiteralInt)
		return lexBegin
	case nextRune == quote:
		return lexStringLiteral
	case nextRune == equals:
		l.readRune()
		l.emit(TokenAssign)
		return lexBegin
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
