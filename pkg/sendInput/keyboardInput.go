package sendInput

type kInput struct {
	inputType inputType
	ki        KeyboardInput
	// Padding is required to make the struct size 40 bytes.
	padding [8]byte
}

// See: https://docs.microsoft.com/en-us/windows/win32/api/winuser/ns-winuser-keybdinput
type KeyboardInput struct {
	WVk         uint16
	WScan       uint16
	DwFlags     KbDwFlag
	Time        uint32
	DwExtraInfo uintptr
}

type KbDwFlag uint32

const (
	ExtendedKey KbDwFlag = 0x0001
	KeyUp       KbDwFlag = 0x0002
	Unicode     KbDwFlag = 0x0004
	ScanCode    KbDwFlag = 0x0008
)
