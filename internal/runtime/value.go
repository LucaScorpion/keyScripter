package runtime

import "fmt"

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

	// Function-specific
	paramKinds []Kind
	variadic   bool
	native     bool
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

func NewFunctionValue(fn RuntimeFunction, paramKinds []Kind) ConcreteValue {
	return ConcreteValue{
		kind:       FunctionKind,
		value:      fn,
		paramKinds: paramKinds,
	}
}

func NewNativeFunctionValue(fn callable, paramKinds []Kind, variadic bool) ConcreteValue {
	return ConcreteValue{
		kind:       FunctionKind,
		value:      fn,
		paramKinds: paramKinds,
		variadic:   variadic,
		native:     true,
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

func (v ConcreteValue) assertKind(expected Kind) {
	if v.kind != expected {
		panic(fmt.Errorf("expected %s, got %s", expected, v.kind))
	}
}

func (v ConcreteValue) Kind() Kind {
	return v.kind
}

func (v ConcreteValue) ParamKinds() []Kind {
	v.assertKind(FunctionKind)
	return v.paramKinds
}

func (v ConcreteValue) ParamKind(index int) Kind {
	v.assertKind(FunctionKind)
	clampedI := index
	if index >= len(v.paramKinds) && v.variadic {
		clampedI = len(v.paramKinds) - 1
	}
	return v.paramKinds[clampedI]
}

func (v ConcreteValue) Variadic() bool {
	v.assertKind(FunctionKind)
	return v.variadic
}

func (v ConcreteValue) call(args []Value, ctx *Context) {
	v.assertKind(FunctionKind)
	v.value.(callable).call(args, ctx)
}

// TODO: Merge VariableValue into ConcreteValue

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
