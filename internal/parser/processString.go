package parser

import (
	"fmt"
)

var escapeSequences = map[rune]rune{
	'n':  '\n',
	't':  '\t',
	'"':  '"',
	'\\': '\\',
}

func processString(raw string) (string, error) {
	result := ""
	escaped := false

	// Go from 1 to len - 2 to remove the wrapping quotes.
	for _, r := range raw[1 : len(raw)-1] {
		switch {
		case escaped:
			seq, ok := escapeSequences[r]
			if !ok {
				return "", fmt.Errorf("invalid escape sequence: \\%s", string(r))
			}

			result += string(seq)
			escaped = false
		case r == '\\':
			escaped = true
		default:
			result += string(r)
		}
	}

	return result, nil
}
