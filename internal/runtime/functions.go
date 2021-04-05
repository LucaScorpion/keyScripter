package runtime

type fnArgs []Value

type RuntimeFn func()

type validateFn func(fnArgs) (RuntimeFn, error)

var Functions = map[string]validateFn{
	"print": printFn,
	"sleep": sleepFn,
}
