package runtime

import "time"

func sleepFn(ms int) {
	time.Sleep(time.Duration(ms) * time.Millisecond)
}
