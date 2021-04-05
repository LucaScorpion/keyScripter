package lexer

import (
	"unicode"
)

func lexBegin(l *lexer) lexFn {
	// Discard leading whitespaces.
	l.readSpace()
	l.discard()

	nextRune := l.peekRune()

	if nextRune == eof {
		l.emit(TokenEOF)
		return nil
	} else if nextRune == newline {
		l.readRune()
		l.emit(TokenNewline)
		return lexBegin
	} else if nextRune == commentStart {
		l.readLine()
		l.emit(TokenComment)
		return lexBegin
	} else if unicode.IsLetter(nextRune) || nextRune == '_' {
		l.readAlphaNum()
		l.emit(TokenIdentifier)
		return lexBegin
	} else if unicode.IsNumber(nextRune) {
		l.readWhile(unicode.IsNumber)
		l.emit(TokenLiteralInt)
		return lexBegin
	} else if nextRune == quote {
		return lexStringLiteral
	} else if nextRune == equals {
		l.readRune()
		l.emit(TokenAssign)
		return lexBegin
	}

	return l.errorf("unexpected character: '%s'", string(nextRune))
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
