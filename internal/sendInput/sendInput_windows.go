package sendInput

import (
	"golang.org/x/sys/windows"
	"unsafe"
)

// See: https://docs.microsoft.com/en-us/windows/win32/api/winuser/ns-winuser-input
var (
	user32        = windows.NewLazySystemDLL("user32.dll")
	sendInputProc = user32.NewProc("SendInput")
)

type inputType uint32

const (
	INPUT_MOUSE    inputType = 0
	INPUT_KEYBOARD inputType = 1
	INPUT_HARDWARE inputType = 2
)

func Keyboard(ki *KeyboardInput) error {
	i := kInput{
		inputType: INPUT_KEYBOARD,
		ki:        *ki,
	}
	return sendRawInput(unsafe.Pointer(&i), unsafe.Sizeof(i))
}

func Mouse(mi *MouseInput) error {
	i := mInput{
		inputType: INPUT_MOUSE,
		mi:        *mi,
	}
	return sendRawInput(unsafe.Pointer(&i), unsafe.Sizeof(i))
}

func sendRawInput(pInputs unsafe.Pointer, cbSize uintptr) error {
	sent, _, err := sendInputProc.Call(
		uintptr(1),       // Number of inputs
		uintptr(pInputs), // Pointer to input
		cbSize,           // Size of an input struct
	)

	if sent != 1 {
		return err
	}
	return nil
}
