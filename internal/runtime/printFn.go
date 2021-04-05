package runtime

import "fmt"

func printFn(format string, values ...interface{}) {
	fmt.Printf(format+"\n", values...)
}
