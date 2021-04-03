package keyScripter

import (
	"fmt"
	"log"
	"syscall"
	"time"
	"unsafe"
)

// Source: https://stackoverflow.com/questions/14489013/simulate-python-keypresses-for-controlling-a-game

var (
	user32        = syscall.NewLazyDLL("user32.dll")
	sendInputProc = user32.NewProc("SendInput")
)

type keyboardInput struct {
	wVk         uint16
	wScan       uint16
	dwFlags     uint32
	time        uint32
	dwExtraInfo uint64
}

type input struct {
	inputType uint32
	ki        keyboardInput
	padding   uint64
}

func Press(scan uint16) {
	var i input
	i.inputType = 1 //INPUT_KEYBOARD
	//i.ki.wVk = 0x0D // virtual key code for enter
	i.ki.wScan = scan
	i.ki.dwFlags = 0x0008

	ret, _, err := sendInputProc.Call(
		uintptr(1),
		uintptr(unsafe.Pointer(&i)),
		unsafe.Sizeof(i),
	)
	log.Printf("ret: %v error: %v", ret, err)
}

func Release(scan uint16) {
	var i input
	i.inputType = 1 //INPUT_KEYBOARD
	//i.ki.wVk = 0x0D // virtual key code for enter
	i.ki.wScan = scan
	i.ki.dwFlags = 0x0008 | 0x0002

	ret, _, err := sendInputProc.Call(
		uintptr(1),
		uintptr(unsafe.Pointer(&i)),
		unsafe.Sizeof(i),
	)
	log.Printf("ret: %v error: %v", ret, err)
}

func PressRelease(scan uint16) {
	fmt.Println(scan)
	Press(scan)
	time.Sleep(20 * time.Millisecond)
	Release(scan)
}
