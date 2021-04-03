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
		l.emit(tokenEOF)
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
	l.emit(tokenComment)
	return lexBegin
}

func lexFunction(l *lexer) lexFn {
	// Function name
	l.readAlphaNum()
	l.emit(tokenIdentifier)

	// Next should be a space, eof, or start of a comment.
	nextRune := l.peekRune()
	if !unicode.IsSpace(nextRune) && nextRune != eof && !strings.HasPrefix(l.currentInput(), commentStart) {
		l.errorf("unexpected character: %s", string(nextRune))
		return nil
	}

	// TODO: arguments

	return lexBegin
}
