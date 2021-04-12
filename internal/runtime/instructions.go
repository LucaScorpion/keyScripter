package runtime

import (
	"time"
)

type Instruction interface {
	Execute(ctx *Context)
}

type FunctionCall struct {
	fn   callable
	args []*Value
}

func NewFunctionCall(fn callable, args []*Value) FunctionCall {
	return FunctionCall{
		fn:   fn,
		args: args,
	}
}

func (f FunctionCall) Execute(ctx *Context) {
	f.fn.call(f.args, ctx)
}

type Assignment struct {
	name string
	val  *Value
}

func NewAssignment(name string, val *Value) Assignment {
	return Assignment{
		name: name,
		val:  val,
	}
}

func (a Assignment) Execute(ctx *Context) {
	ctx.SetValue(a.name, a.val.Resolve(ctx))
}

type TimestampInstr struct {
	ms          int64
	instruction Instruction
}

func NewTimestampInstr(ms int64, instruction Instruction) TimestampInstr {
	return TimestampInstr{
		ms:          ms,
		instruction: instruction,
	}
}

type Timestamps struct {
	instructions []TimestampInstr
}

func NewTimestamps(instructions []TimestampInstr) Timestamps {
	return Timestamps{
		instructions: instructions,
	}
}

func unixMillis() int64 {
	return time.Now().UnixNano() / 1_000_000
}

func (t Timestamps) Execute(ctx *Context) {
	startMillis := unixMillis()
	for _, instr := range t.instructions {
		// Get the instruction execution timestamp, wait for it.
		execAtMillis := startMillis + instr.ms
		waitMillis := execAtMillis - unixMillis()
		if waitMillis > 0 {
			time.Sleep(time.Duration(waitMillis) * time.Millisecond)
		}

		instr.instruction.Execute(ctx)
	}
}
