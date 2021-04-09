package runtime

type Kind string

const (
	StringKind   Kind = "string"
	NumberKind   Kind = "number"
	AnyKind      Kind = "any"
	FunctionKind Kind = "function"
)

type Value interface {
	Resolve(ctx *Context) ConcreteValue
}

type ConcreteValue struct {
	kind  Kind
	value interface{}
}

func NewStringValue(val string) ConcreteValue {
	return ConcreteValue{
		kind:  StringKind,
		value: val,
	}
}

func NewNumberValue(val int) ConcreteValue {
	return ConcreteValue{
		kind:  NumberKind,
		value: val,
	}
}

func NewFunctionValue(val []Instruction) ConcreteValue {
	return ConcreteValue{
		kind: FunctionKind,
		value: func(ctx *Context) {
			funcCtx := NewContext(ctx)
			for _, i := range val {
				i.Execute(funcCtx)
			}
		},
	}
}

func NewEmptyValue(kind Kind) ConcreteValue {
	return ConcreteValue{
		kind: kind,
	}
}

func (v ConcreteValue) Resolve(_ *Context) ConcreteValue {
	return v
}

func (v ConcreteValue) Kind() Kind {
	return v.kind
}

type VariableValue struct {
	ref string
}

func NewVariableValue(ref string) VariableValue {
	return VariableValue{
		ref: ref,
	}
}

func (v VariableValue) Ref() string {
	return v.ref
}

func (v VariableValue) Resolve(ctx *Context) ConcreteValue {
	return ctx.GetValue(v.ref)
}
