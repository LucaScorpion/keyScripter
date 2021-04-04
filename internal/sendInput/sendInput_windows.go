package sendInput

import (
	"fmt"
	"log"
	"reflect"
	"syscall"
	"time"
	"unsafe"
)

var (
	user32        = syscall.NewLazyDLL("user32.dll")
	sendInputProc = user32.NewProc("SendInput")
)

type keyboardInput struct {
	// See: https://docs.microsoft.com/en-us/windows/win32/inputdev/virtual-key-codes
	wVk         uint16
	wScan       uint16
	dwFlags     dwFlag
	time        uint32
	dwExtraInfo uint64
}

type input struct {
	inputType inputType
	ki        keyboardInput
	padding   uint64
}

// See: https://docs.microsoft.com/en-us/windows/win32/api/winuser/ns-winuser-input
type inputType uint32

const (
	INPUT_MOUSE    = 0
	INPUT_KEYBOARD = 1
	INPUT_HARDWARE = 2
)

// See: https://docs.microsoft.com/en-us/windows/win32/api/winuser/ns-winuser-keybdinput
type dwFlag uint32

const (
	KEYEVENTF_KEYUP    dwFlag = 0x0002
	KEYEVENTF_SCANCODE dwFlag = 0x0008
)

func Test() {
	var i = []input{
		{
			inputType: INPUT_KEYBOARD,
			ki: keyboardInput{
				wVk: 0x5B,
			},
		},
		{
			inputType: INPUT_KEYBOARD,
			ki: keyboardInput{
				wVk: 0x52,
			},
		},
		{
			inputType: INPUT_KEYBOARD,
			ki: keyboardInput{
				wVk:     0x52,
				dwFlags: KEYEVENTF_KEYUP,
			},
		},
		{
			inputType: INPUT_KEYBOARD,
			ki: keyboardInput{
				wVk:     0x5B,
				dwFlags: KEYEVENTF_KEYUP,
			},
		},
	}

	err := sendInput(i...)
	if err != nil {
		panic(err)
	}
}

func sendInput(i ...input) error {
	if len(i) == 0 {
		return nil
	}

	sentPtr, _, err := sendInputProc.Call(
		uintptr(len(i)),                 // Number of inputs
		uintptr(unsafe.Pointer(&i[0])),  // Pointer to inputs
		reflect.TypeOf(i).Elem().Size(), // Size of an input struct
	)

	sent := int(sentPtr)
	if sent != len(i) {
		return err
	}
	return nil
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
