package sendInput

type mouseInput struct {
	dx          int64
	dy          int64
	mouseData   uint32
	dwFlags     mDwFlag
	time        uint32
	dwExtraInfo uintptr
}

type mDwFlag uint32

const (
	MOUSEEVENTF_MOVE            mDwFlag = 0x0001
	MOUSEEVENTF_LEFTDOWN        mDwFlag = 0x0002
	MOUSEEVENTF_LEFTUP          mDwFlag = 0x0004
	MOUSEEVENTF_RIGHTDOWN       mDwFlag = 0x0008
	MOUSEEVENTF_RIGHTUP         mDwFlag = 0x0010
	MOUSEEVENTF_MIDDLEDOWN      mDwFlag = 0x0020
	MOUSEEVENTF_MIDDLEUP        mDwFlag = 0x0040
	MOUSEEVENTF_XDOWN           mDwFlag = 0x0080
	MOUSEEVENTF_XUP             mDwFlag = 0x0100
	MOUSEEVENTF_WHEEL           mDwFlag = 0x0800
	MOUSEEVENTF_HWHEEL          mDwFlag = 0x1000
	MOUSEEVENTF_MOVE_NOCOALESCE mDwFlag = 0x2000
	MOUSEEVENTF_VIRTUALDESK     mDwFlag = 0x4000
	MOUSEEVENTF_ABSOLUTE        mDwFlag = 0x8000
)
