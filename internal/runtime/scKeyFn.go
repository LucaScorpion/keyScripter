package runtime

import "github.com/LucaScorpion/keyScripter/pkg/sendInput"

func scKeyDown(key int) {
	err := sendInput.Keyboard(&sendInput.KeyboardInput{
		WScan:   uint16(key),
		DwFlags: sendInput.ScanCode,
	})
	if err != nil {
		panic(err)
	}
}

func scKeyUp(key int) {
	err := sendInput.Keyboard(&sendInput.KeyboardInput{
		WScan:   uint16(key),
		DwFlags: sendInput.ScanCode | sendInput.KeyUp,
	})
	if err != nil {
		panic(err)
	}
}
