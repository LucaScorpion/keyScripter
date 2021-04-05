package runtime

type Instruction interface {
	Execute(ctx *context)
}

type FunctionCall struct {
	Fn   *ScriptFn
	Args []Value
}

func (f FunctionCall) Execute(ctx *context) {
	f.Fn.call(f.Args, ctx)
}

type Assignment struct {
	Name string
	Val  Value
}

func (a Assignment) Execute(ctx *context) {
	ctx.setValue(a.Name, a.Val)
}
