package runtime

import "fmt"

func printFn(args fnArgs) (RuntimeFn, error) {
	if len(args) < 1 {
		return nil, fmt.Errorf("print requires at least 1 argument")
	}

	format, ok := args[0].(string)
	if !ok {
		return nil, fmt.Errorf("the first argument of print must be a string")
	}

	return func() {
		fmt.Printf(format+"\n", args[1:]...)
	}, nil
}
