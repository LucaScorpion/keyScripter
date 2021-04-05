package runtime

import (
	"fmt"
	"reflect"
)

var Functions = map[string]*ScriptFn{
	"print": newScriptFn(printFn),
	"sleep": newScriptFn(sleepFn),
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
		res.params = append(res.params, kindFromType(t.In(i)))
	}

	return res
}

func (fn *ScriptFn) Validate(args []Kind) error {
	// Check if the argument count matches.
	if len(args) < len(fn.params) || (!fn.variadic && len(args) > len(fn.params)) {
		return fmt.Errorf("mismatched argument count, expected %d but got %d", len(fn.params), len(args))
	}

	// Check if the argument types match.
	for argI := 0; argI < len(args); argI++ {
		// Get the param index, to account for variadic parameters.
		paramI := argI
		if argI > len(fn.params) {
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

func (fn *ScriptFn) call(args []Value, ctx *context) {
	in := make([]reflect.Value, len(args))
	for i := 0; i < len(args); i++ {
		in = append(in, reflect.ValueOf(args[i].Value(ctx)))
	}

	fn.rawFn.Call(in)
}
