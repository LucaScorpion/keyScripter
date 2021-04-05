package runtime

import (
	"fmt"
	"reflect"
)

type Value interface {
	Kind(ctx *context) Kind
	Value(ctx *context) interface{}
}

type Kind string

const (
	StringKind Kind = "string"
	NumberKind Kind = "number"
	AnyKind    Kind = "any"
)

var kindMap = map[reflect.Kind]Kind{
	reflect.String:    StringKind,
	reflect.Int:       NumberKind,
	reflect.Interface: AnyKind,
}

func kindFromType(t reflect.Type) Kind {
	k := t.Kind()
	if r, ok := kindMap[k]; ok {
		return r
	}

	if k == reflect.Slice {
		// Don't recurse here, since we can only handle 1-deep slice types.
		if r, ok := kindMap[t.Elem().Kind()]; ok {
			return r
		}
	}

	panic(fmt.Errorf("invalid value kind: %s", k.String()))
}

type StringValue struct {
	val string
}

func NewStringValue(val string) StringValue {
	return StringValue{val: val}
}

func (v StringValue) Kind(_ *context) Kind {
	return StringKind
}

func (v StringValue) Value(_ *context) interface{} {
	return v.val
}

type NumberValue struct {
	val int
}

func NewNumberValue(val int) NumberValue {
	return NumberValue{val: val}
}

func (v NumberValue) Kind(_ *context) Kind {
	return NumberKind
}

func (v NumberValue) Value(_ *context) interface{} {
	return v.val
}

type VariableValue struct {
	ref string
}

func NewVariableValue(ref string) VariableValue {
	return VariableValue{ref: ref}
}

func (v VariableValue) Kind(ctx *context) Kind {
	return ctx.getValue(v.ref).Kind(ctx)
}

func (v VariableValue) Value(ctx *context) interface{} {
	return ctx.getValue(v.ref).Value(ctx)
}
