package parser

import (
	"bufio"
	"fmt"
	"io"
	"strings"
)

type Parser struct {
	rd  *bufio.Reader
	pos Position
}

type Position struct {
	line int
	col  int
}

func NewParser(input io.Reader) *Parser {
	return &Parser{
		rd: bufio.NewReader(input),
		pos: Position{
			line: 1,
			col:  1,
		},
	}
}

func (p *Parser) Parse() error {
	var err error
	var instructions []*Instruction

	for err == nil {
		var next *Instruction
		next, err = p.nextInstruction()

		// Store the instruction.
		if next != nil {
			instructions = append(instructions, next)
			fmt.Println(next)
		}
	}

	if err == io.EOF {
		return nil
	}
	return err
}

func (p *Parser) nextInstruction() (*Instruction, error) {
	line, err := p.readLine()

	// Check for empty lines.
	if len(line) == 0 {
		return nil, err
	}

	var iType InstructionType = functionCall

	// If the line starts with a '#', it is a comment.
	if line[0] == '#' {
		iType = comment
	}

	return &Instruction{
		iType: iType,
		raw:   line,
	}, err
}

func (p *Parser) readLine() (string, error) {
	str, err := p.rd.ReadString('\n')

	// Trim any trailing \r and \n.
	str = strings.TrimRight(str, "\r\n")

	p.pos.line += 1
	p.pos.col = 1
	return str, err
}
