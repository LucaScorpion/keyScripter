package keyScripter

import (
	"github.com/micmonay/keybd_event"
	"runtime"
	"time"
)

func Press() {
	kb, err := keybd_event.NewKeyBonding()
	if err != nil {
		panic(err)
	}

	// According to the readme, on Linux we need to wait 2 seconds.
	if runtime.GOOS == "linux" {
		time.Sleep(2 * time.Second)
	}

	kb.SetKeys(keybd_event.VK_R)
	kb.HasSuper(true)
	kb.Launching()
}
