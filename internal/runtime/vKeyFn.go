package runtime

import "github.com/LucaScorpion/keyScripter/pkg/sendInput"

func vKeyDown(key int) {
	err := sendInput.Keyboard(&sendInput.KeyboardInput{
		WVk: uint16(key),
	})
	if err != nil {
		panic(err)
	}
}

func vKeyUp(key int) {
	err := sendInput.Keyboard(&sendInput.KeyboardInput{
		WVk:     uint16(key),
		DwFlags: sendInput.KeyUp,
	})
	if err != nil {
		panic(err)
	}
}

func vKeyPress(key int) {
	vKeyDown(key)
	vKeyUp(key)
}
