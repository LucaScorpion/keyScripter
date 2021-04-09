package runtime

import (
	"fmt"
	"reflect"
)

var Functions = map[string]*ScriptFn{
	"print":     newScriptFn(printFn),
	"sleep":     newScriptFn(sleepFn),
	"pause":     newScriptFn(pauseFn),
	"vKeyDown":  newScriptFn(vKeyDown),
	"vKeyUp":    newScriptFn(vKeyUp),
	"vKeyPress": newScriptFn(vKeyPress),
}

type ScriptFn struct {
	rawFn    reflect.Value
	params   []Kind
	variadic bool
}

func newScriptFn(fn interface{}) *ScriptFn {
	t := reflect.TypeOf(fn)
	if t.Kind() != reflect.Func {
		panic(fmt.Errorf("cannot make runtime function from %s", t.Kind()))
	}

	res := &ScriptFn{
		rawFn:    reflect.ValueOf(fn),
		params:   make([]Kind, t.NumIn()),
		variadic: t.IsVariadic(),
	}

	for i := 0; i < t.NumIn(); i++ {
		res.params[i] = kindFromType(t.In(i))
	}

	return res
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

func (fn *ScriptFn) Validate(args []Kind) error {
	// Check if the argument count matches.
	if (fn.variadic && len(args) < len(fn.params)-1) || (!fn.variadic && len(args) != len(fn.params)) {
		return fmt.Errorf("mismatched argument count, expected %d but got %d", len(fn.params), len(args))
	}

	// Check if the argument types match.
	for argI := 0; argI < len(args); argI++ {
		// Get the param index, to account for variadic parameters.
		paramI := argI
		if argI > len(fn.params)-1 {
			paramI = len(fn.params) - 1
		}

		paramKind := fn.params[paramI]
		argKind := args[argI]
		if paramKind != AnyKind && paramKind != argKind {
			return fmt.Errorf("mismatched argument type, expected %s but got %s", paramKind, argKind)
		}
	}

	return nil
}

func (fn *ScriptFn) call(args []Value, ctx *Context) {
	in := make([]reflect.Value, len(args))
	for i := 0; i < len(args); i++ {
		in[i] = reflect.ValueOf(args[i].Resolve(ctx).value)
	}

	fn.rawFn.Call(in)
}
