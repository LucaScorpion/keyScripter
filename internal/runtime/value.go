package runtime

type Value interface {
	kind() string
	resolveKind( /* TODO: context */ ) string
	resolveValue( /* TODO: context */ ) interface{}
}

const (
	kindString   = "string"
	kindNumber   = "number"
	kindVariable = "variable"
)

type StringValue struct {
	Val string
}

func (v StringValue) kind() string {
	return kindString
}

func (v StringValue) resolveKind() string {
	return v.kind()
}

func (v StringValue) resolveValue() interface{} {
	return v.Val
}

type NumberValue struct {
	Val int
}

func (v NumberValue) kind() string {
	return kindNumber
}

func (v NumberValue) resolveKind() string {
	return v.kind()
}

func (v NumberValue) resolveValue() interface{} {
	return v.Val
}

type VariableValue struct {
	VarName string
}

func (v VariableValue) kind() string {
	return kindVariable
}

func (v VariableValue) resolveKind() string {
	// TODO
	return "todo"
}

func (v VariableValue) resolveValue() interface{} {
	// TODO
	return "todo"
}
