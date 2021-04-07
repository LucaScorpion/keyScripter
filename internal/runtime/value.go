package runtime

type Kind string

const (
	StringKind Kind = "string"
	NumberKind Kind = "number"
	AnyKind    Kind = "any"
)

type Value interface {
	Resolve(ctx *Context) ConcreteValue
}

type ConcreteValue struct {
	Kind  Kind
	Value interface{}
}

func NewStringValue(val string) ConcreteValue {
	return ConcreteValue{
		Kind:  StringKind,
		Value: val,
	}
}

func NewNumberValue(val int) ConcreteValue {
	return ConcreteValue{
		Kind:  NumberKind,
		Value: val,
	}
}

func (v ConcreteValue) Resolve(_ *Context) ConcreteValue {
	return v
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
