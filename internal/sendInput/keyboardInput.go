package sendInput

type keyboardInput struct {
	wVk         uint16
	wScan       uint16
	dwFlags     kbDwFlag
	time        uint32
	dwExtraInfo uintptr
}

type kbDwFlag uint32

const (
	KEYEVENTF_EXTENDEDKEY kbDwFlag = 0x0001
	KEYEVENTF_KEYUP       kbDwFlag = 0x0002
	KEYEVENTF_UNICODE     kbDwFlag = 0x0004
	KEYEVENTF_SCANCODE    kbDwFlag = 0x0008
)
