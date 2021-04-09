package runtime

type Instruction interface {
	Execute(ctx *Context)
}

type FunctionCall struct {
	fn   callable
	args []Value
}

func NewFunctionCall(fn callable, args []Value) FunctionCall {
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
	val  Value
}

func NewAssignment(name string, val Value) Assignment {
	return Assignment{
		name: name,
		val:  val,
	}
}

func (a Assignment) Execute(ctx *Context) {
	ctx.SetValue(a.name, a.val.Resolve(ctx))
}
