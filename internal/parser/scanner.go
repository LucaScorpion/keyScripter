package parser

import (
	"bufio"
	"io"
	"strings"
)

type scanner struct {
	rd  *bufio.Reader
	pos *position
}

type position struct {
	line int
	col  int
}

func newScanner(r io.Reader) *scanner {
	return &scanner{
		rd: bufio.NewReader(r),
		pos: &position{
			line: 1,
			col:  1,
		},
	}
}

func (s *scanner) readLine() (string, error) {
	str, err := s.rd.ReadString('\n')

	// Trim any trailing \r and \n.
	str = strings.TrimRight(str, "\r\n")

	s.pos.line += 1
	s.pos.col = 1
	return str, err
}
