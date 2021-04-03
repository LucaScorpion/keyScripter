package parser

import (
	"fmt"
	"github.com/LucaScorpion/tas-scripter/internal/lexer"
)

func Parse(input string) {
	l := lexer.NewLexer(input)
	go l.Run()
	for t := range l.Tokens {
		fmt.Println(t)
	}
}
