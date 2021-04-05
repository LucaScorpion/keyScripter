package runtime

import (
	"fmt"
	"time"
)

func sleepFn(args fnArgs) (RuntimeFn, error) {
	if len(args) != 1 {
		return nil, fmt.Errorf("sleep requires 1 argument")
	}

	argKind := args[0].resolveKind()
	if argKind != kindNumber {
		return nil, fmt.Errorf("the first argument of sleep must be a number, got: %s", argKind)
	}

	return func() {
		time.Sleep(time.Duration(args[0].resolveValue().(int)) * time.Millisecond)
	}, nil
}
