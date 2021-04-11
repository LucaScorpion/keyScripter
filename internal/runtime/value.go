package runtime

import "fmt"

type Kind string

const (
	StringKind   Kind = "string"
	NumberKind   Kind = "number"
	AnyKind      Kind = "any"
	FunctionKind Kind = "function"
	VariableKind Kind = "variable"
)

type Value struct {
	kind  Kind
	value interface{}

	// Function
	paramKinds []Kind
	variadic   bool
	native     bool
}

func NewStringValue(val string) *Value {
	return &Value{
		kind:  StringKind,
		value: val,
	}
}

func NewNumberValue(val int) *Value {
	return &Value{
		kind:  NumberKind,
		value: val,
	}
}

func NewFunctionValue(fn *RuntimeFunction, paramKinds []Kind) *Value {
	return &Value{
		kind:       FunctionKind,
		value:      fn,
		paramKinds: paramKinds,
	}
}

func NewNativeFunctionValue(fn callable, paramKinds []Kind, variadic bool) *Value {
	return &Value{
		kind:       FunctionKind,
		value:      fn,
		paramKinds: paramKinds,
		variadic:   variadic,
		native:     true,
	}
}

func NewEmptyValue(kind Kind) *Value {
	return &Value{
		kind: kind,
	}
}

func NewVariableValue(ref string) *Value {
	return &Value{
		kind:  VariableKind,
		value: ref,
	}
}

func (v *Value) Resolve(ctx *Context) *Value {
	if v.kind == VariableKind {
		return ctx.GetValue(v.value.(string))
	}
	return v
}

func (v *Value) assertKind(expected Kind) {
	if v.kind != expected {
		panic(fmt.Errorf("expected %s, got %s", expected, v.kind))
	}
}

func (v *Value) RawValue() interface{} {
	return v.value
}

func (v *Value) Kind() Kind {
	return v.kind
}

func (v *Value) ParamKinds() []Kind {
	v.assertKind(FunctionKind)
	return v.paramKinds
}

func (v *Value) ParamKind(index int) Kind {
	v.assertKind(FunctionKind)
	clampedI := index
	if index >= len(v.paramKinds) && v.variadic {
		clampedI = len(v.paramKinds) - 1
	}
	return v.paramKinds[clampedI]
}

func (v *Value) Variadic() bool {
	v.assertKind(FunctionKind)
	return v.variadic
}

func (v *Value) call(args []*Value, ctx *Context) {
	v.assertKind(FunctionKind)
	v.value.(callable).call(args, ctx)
}
