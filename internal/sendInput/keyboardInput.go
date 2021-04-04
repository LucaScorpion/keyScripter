package sendInput

type kInput struct {
	inputType inputType
	ki        KeyboardInput
	padding   [8]byte
}

type KeyboardInput struct {
	WVk         uint16
	WScan       uint16
	DwFlags     KbDwFlag
	Time        uint32
	DwExtraInfo uintptr
}

type KbDwFlag uint32

const (
	KEYEVENTF_EXTENDEDKEY KbDwFlag = 0x0001
	KEYEVENTF_KEYUP       KbDwFlag = 0x0002
	KEYEVENTF_UNICODE     KbDwFlag = 0x0004
	KEYEVENTF_SCANCODE    KbDwFlag = 0x0008
)
