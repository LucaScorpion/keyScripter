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
	return VariableValue{ref: ref}
}

func (v VariableValue) Resolve(ctx *Context) ConcreteValue {
	return ctx.GetValue(v.ref)
}
