package sendInput

import (
	"golang.org/x/sys/windows"
	"unsafe"
)

var (
	user32 = windows.NewLazySystemDLL("user32.dll")
	// See: https://docs.microsoft.com/en-us/windows/win32/api/winuser/nf-winuser-sendinput
	sendInputProc = user32.NewProc("SendInput")
)

type inputType uint32

const (
	mouse    inputType = 0
	keyboard inputType = 1
)

// Send a KeyboardInput to the system.
func Keyboard(ki *KeyboardInput) error {
	i := kInput{
		inputType: keyboard,
		ki:        *ki,
	}
	return sendInput(1, unsafe.Pointer(&i), unsafe.Sizeof(i))
}

// Send a MouseInput to the system.
func Mouse(mi *MouseInput) error {
	i := mInput{
		inputType: mouse,
		mi:        *mi,
	}
	return sendInput(1, unsafe.Pointer(&i), unsafe.Sizeof(i))
}

// Call sendInputProc with the given pInputs.
// cInputs is the number of inputs in the pInputs array.
// pInputs is a pointer to the first input.
// cbSize is the size of a single input struct, as a result of unsafe.Sizeof.
// This is constant and should always be 40 bytes,
// but is passed here to ensure the structs are formed correctly.
func sendInput(cInputs uintptr, pInputs unsafe.Pointer, cbSize uintptr) error {
	sent, _, err := sendInputProc.Call(cInputs, uintptr(pInputs), cbSize)

	// The return value is the number of successfully inserted events,
	// and the returned error is always non-nil.
	if sent != cInputs {
		return err
	}
	return nil
}
