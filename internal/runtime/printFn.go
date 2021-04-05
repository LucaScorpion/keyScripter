package runtime

import "fmt"

func printFn(args fnArgs) (RuntimeFn, error) {
	if len(args) < 1 {
		return nil, fmt.Errorf("print requires at least 1 argument")
	}

	formatKind := args[0].resolveKind()
	if formatKind != kindString {
		return nil, fmt.Errorf("the first argument of print must be a string, got: %s", formatKind)
	}

	return func() {
		// TODO
		fmt.Printf(args[0].resolveValue().(string) + "\n" /*, args[1:]...*/)
	}, nil
}
