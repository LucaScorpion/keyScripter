package runtime

import (
	"bufio"
	"fmt"
	"os"
)

func pauseFn() {
	fmt.Print("Press enter to continue...")
	bufio.NewReader(os.Stdin).ReadLine()
}
