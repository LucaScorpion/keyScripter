package sendInput

type mInput struct {
	inputType inputType
	mi        MouseInput
}

type MouseInput struct {
	Dx          int32
	Dy          int32
	MouseData   uint32
	DwFlags     MDwFlag
	Time        uint32
	DwExtraInfo uintptr
}

type MDwFlag uint32

const (
	MOUSEEVENTF_MOVE            MDwFlag = 0x0001
	MOUSEEVENTF_LEFTDOWN        MDwFlag = 0x0002
	MOUSEEVENTF_LEFTUP          MDwFlag = 0x0004
	MOUSEEVENTF_RIGHTDOWN       MDwFlag = 0x0008
	MOUSEEVENTF_RIGHTUP         MDwFlag = 0x0010
	MOUSEEVENTF_MIDDLEDOWN      MDwFlag = 0x0020
	MOUSEEVENTF_MIDDLEUP        MDwFlag = 0x0040
	MOUSEEVENTF_XDOWN           MDwFlag = 0x0080
	MOUSEEVENTF_XUP             MDwFlag = 0x0100
	MOUSEEVENTF_WHEEL           MDwFlag = 0x0800
	MOUSEEVENTF_HWHEEL          MDwFlag = 0x1000
	MOUSEEVENTF_MOVE_NOCOALESCE MDwFlag = 0x2000
	MOUSEEVENTF_VIRTUALDESK     MDwFlag = 0x4000
	MOUSEEVENTF_ABSOLUTE        MDwFlag = 0x8000
)
