package parser

import "fmt"

type InstructionType string

const (
	comment      InstructionType = "comment"
	functionCall InstructionType = "functionCall"
)

type Instruction struct {
	iType InstructionType
	raw   string
}

func (i *Instruction) String() string {
	return fmt.Sprintf("%s: %s", i.iType, i.raw)
}
