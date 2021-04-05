package runtime

type Instruction interface {
	Execute(ctx *Context)
}

type FunctionCall struct {
	Fn   *ScriptFn
	Args []Value
}

func (f FunctionCall) Execute(ctx *Context) {
	f.Fn.call(f.Args, ctx)
}

type Assignment struct {
	Name string
	Val  Value
}

func (a Assignment) Execute(ctx *Context) {
	ctx.SetValue(a.Name, a.Val)
}
