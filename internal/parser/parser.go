package parser

import (
	"fmt"
	"io"
)

type Parser struct {
	scan *scanner
}

func NewParser(r io.Reader) *Parser {
	return &Parser{
		scan: newScanner(r),
	}
}

func (p *Parser) Parse() error {
	var err error
	var instructions []*Instruction
	prevPos := &position{
		line: 1,
		col:  1,
	}

	for err == nil {
		// Store the previous position.
		prevPos = &position{
			line: p.scan.pos.line,
			col:  p.scan.pos.col,
		}

		// Parse the next instruction.
		var next *Instruction
		next, err = p.nextInstruction()

		// Store the instruction.
		if next != nil {
			instructions = append(instructions, next)
		}
	}

	if err == io.EOF {
		return nil
	}
	return fmt.Errorf("an error occurred while parsing the script (line %d, col %d): %s", prevPos.line, prevPos.col, err)
}

func (p *Parser) nextInstruction() (*Instruction, error) {
	line, err := p.scan.readLine()

	// Check for empty lines.
	if len(line) == 0 {
		return nil, err
	}

	var iType = functionCall

	// If the line starts with a '#', it is a comment.
	if line[0] == '#' {
		iType = comment
	}

	return &Instruction{
		iType: iType,
		raw:   line,
	}, err
}
