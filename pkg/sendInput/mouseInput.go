package sendInput

type mInput struct {
	inputType inputType
	mi        MouseInput
}

// See: https://docs.microsoft.com/en-us/windows/win32/api/winuser/ns-winuser-mouseinput
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
	Move           MDwFlag = 0x0001
	LeftDown       MDwFlag = 0x0002
	LeftUp         MDwFlag = 0x0004
	RightDown      MDwFlag = 0x0008
	RightUp        MDwFlag = 0x0010
	MiddleDown     MDwFlag = 0x0020
	MiddleUp       MDwFlag = 0x0040
	XDown          MDwFlag = 0x0080
	XUp            MDwFlag = 0x0100
	Wheel          MDwFlag = 0x0800
	HWheel         MDwFlag = 0x1000
	MoveNoCoalesce MDwFlag = 0x2000
	VirtualDesk    MDwFlag = 0x4000
	Absolute       MDwFlag = 0x8000
)
