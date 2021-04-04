package main

import (
	"fmt"
	"github.com/LucaScorpion/keyScripter/internal/parser"
	"github.com/LucaScorpion/keyScripter/internal/sendInput"
	"io/ioutil"
)

func main() {
	// Input test

	sendInput.Keyboard(&sendInput.KeyboardInput{
		WVk: 0x5B,
	})
	sendInput.Keyboard(&sendInput.KeyboardInput{
		WVk:     0x5B,
		DwFlags: sendInput.KEYEVENTF_KEYUP,
	})

	sendInput.Mouse(&sendInput.MouseInput{
		Dx:      200,
		Dy:      200,
		DwFlags: sendInput.MOUSEEVENTF_MOVE,
	})

	// WScan finder

	//time.Sleep(2 * time.Second)
	//sendInput.PressRelease(28)
	//time.Sleep(1000 * time.Millisecond)
	//sendInput.PressRelease(208) // down
	//time.Sleep(100 * time.Millisecond)
	//keyScripter.PressRelease(200) // up

	//for i := 200; i < 255; i++ {
	//	keyScripter.PressRelease(uint16(i))
	//	time.Sleep(1000 * time.Millisecond)
	//}

	//time.Sleep(1500 * time.Millisecond)
	//keyScripter.PressRelease(50)
	return

	// Script test

	b, _ := ioutil.ReadFile("simple.txt")
	p := parser.NewParser(string(b))

	if err := p.Parse(); err != nil {
		fmt.Printf("An error occurred while parsing the script: %s\n", err)
		return
	}

	if err := p.Prepare(); err != nil {
		fmt.Printf("An error occurred while preparing the script: %s\n", err)
		return
	}

	p.Run()
}
