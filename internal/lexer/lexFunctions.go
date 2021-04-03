package lexer

import (
	"strings"
	"unicode"
)

func lexBegin(l *lexer) lexFn {
	l.readSpace(true)
	l.discard()

	// Check if we are at EOF.
	if l.peekRune() == eof {
		l.emit(TokenEOF)
		return nil
	}

	if strings.HasPrefix(l.currentInput(), commentStart) {
		return lexComment
	} else {
		return lexFunction
	}
}

func lexComment(l *lexer) lexFn {
	l.readLine()
	l.emit(TokenComment)
	return lexBegin
}

func lexFunction(l *lexer) lexFn {
	// Function name
	l.readAlphaNum()
	l.emit(TokenFuncName)

	if !l.expectSpace() {
		return nil
	}
	return lexArguments
}

func lexArguments(l *lexer) lexFn {
	l.readSpace(false)
	l.discard()

	nextRune := l.peekRune()

	// Check if we are at EOF or EOL.
	if nextRune == eof || nextRune == newline {
		return lexBegin
	}

	// String literal
	if nextRune == quote {
		l.readStringLiteral()
		l.emit(TokenLiteralString)
	}

	// Int literal
	if unicode.IsNumber(nextRune) {
		l.readWhile(unicode.IsNumber)
		l.emit(TokenLiteralInt)
	}

	if !l.expectSpace() {
		return nil
	}
	return lexArguments
}
