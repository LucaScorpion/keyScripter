package main

import (
	"github.com/LucaScorpion/tas-scripter/internal/parser"
	"os"
)

func main() {
	f, err := os.Open("example.txt")
	if err != nil {
		panic(err)
	}
	p := parser.NewParser(f)

	err = p.Parse()
	if err != nil {
		panic(err)
	}
}
