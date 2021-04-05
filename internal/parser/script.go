package parser

import "github.com/LucaScorpion/keyScripter/internal/runtime"

type Script struct {
	funcs []*runtime.RuntimeFn
}

func (s *Script) Run() {
	for _, f := range s.funcs {
		(*f)()
	}
}
