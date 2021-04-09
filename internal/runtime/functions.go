package runtime

import (
	"fmt"
	"reflect"
)

var Functions = map[string]ConcreteValue{
	"print":     makeNativeFunction(printFn),
	"sleep":     makeNativeFunction(sleepFn),
	"pause":     makeNativeFunction(pauseFn),
	"vKeyDown":  makeNativeFunction(vKeyDown),
	"vKeyUp":    makeNativeFunction(vKeyUp),
	"vKeyPress": makeNativeFunction(vKeyPress),
}

type callable interface {
	call(args []Value, ctx *Context)
}

type NativeFunction struct {
	rawFn reflect.Value
}

type RuntimeFunction struct {
	paramNames   []string
	instructions []Instruction
}

func makeNativeFunction(fn interface{}) ConcreteValue {
	t := reflect.TypeOf(fn)
	if t.Kind() != reflect.Func {
		panic(fmt.Errorf("cannot make native function from %s", t.Kind()))
	}

	paramKinds := make([]Kind, t.NumIn())
	for i := 0; i < t.NumIn(); i++ {
		paramKinds[i] = kindFromType(t.In(i))
	}

	return NewNativeFunctionValue(NativeFunction{
		rawFn: reflect.ValueOf(fn),
	}, paramKinds, t.IsVariadic())
}

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

func (f NativeFunction) call(args []Value, ctx *Context) {
	in := make([]reflect.Value, len(args))
	for i, a := range args {
		in[i] = reflect.ValueOf(a.Resolve(ctx).value)
	}

	f.rawFn.Call(in)
}

func NewRuntimeFunction(paramNames []string, instructions []Instruction) RuntimeFunction {
	return RuntimeFunction{
		paramNames:   paramNames,
		instructions: instructions,
	}
}

func (f RuntimeFunction) call(args []Value, ctx *Context) {
	// Set all argument values in context.
	funcCtx := NewContext(ctx)
	for i, v := range args {
		funcCtx.SetValue(f.paramNames[i], v.Resolve(ctx))
	}

	// Execute all instructions.
	for _, i := range f.instructions {
		i.Execute(funcCtx)
	}
}
