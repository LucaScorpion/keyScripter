package runtime

type fnArgs []interface{}

type RuntimeFn func()

type validateFn func(fnArgs) (RuntimeFn, error)

var Functions = map[string]validateFn{
	"print": printFn,
	"wait":  waitFn,
}
