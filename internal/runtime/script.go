package runtime

type Script struct {
	instructions []Instruction
}

func NewScript(instructions []Instruction) *Script {
	return &Script{instructions: instructions}
}

func (s *Script) Run() {
	ctx := &context{
		values: make(map[string]Value),
	}
	for _, f := range s.instructions {
		f.Execute(ctx)
	}
}
