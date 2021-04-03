package runtime

import (
	"fmt"
	"time"
)

func waitFn(args fnArgs) (RuntimeFn, error) {
	if len(args) != 1 {
		return nil, fmt.Errorf("wait requires 1 argument")
	}

	t, ok := args[0].(int)
	if !ok {
		return nil, fmt.Errorf("the first argument of wait must be a number")
	}

	return func() {
		time.Sleep(time.Duration(t) * time.Millisecond)
	}, nil
}
