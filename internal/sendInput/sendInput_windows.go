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
	mouse    inputType = 0
	keyboard inputType = 1
)

func Keyboard(ki *KeyboardInput) error {
	i := kInput{
		inputType: keyboard,
		ki:        *ki,
	}
	return sendRawInput(unsafe.Pointer(&i), unsafe.Sizeof(i))
}

func Mouse(mi *MouseInput) error {
	i := mInput{
		inputType: mouse,
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
