package main

import (
	"fmt"
	"github.com/LucaScorpion/tas-scripter/internal/lexer"
	"io/ioutil"
)

func main() {
	b, _ := ioutil.ReadFile("example.txt")
	l := lexer.NewLexer(string(b))
	go l.Run()
	for t := range l.Tokens {
		fmt.Println(t)
	}
}
