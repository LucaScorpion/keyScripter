package main

import (
	"fmt"
	"github.com/LucaScorpion/tas-scripter/internal/parser"
	"io/ioutil"
)

func main() {
	b, _ := ioutil.ReadFile("simple.txt")
	p := parser.NewParser(string(b))

	if err := p.Parse(); err != nil {
		fmt.Printf("An error occurred while parsing the script: %s\n", err)
		return
	}

	if err := p.Prepare(); err != nil {
		fmt.Printf("An error occurred while preparing the script: %s\n", err)
		return
	}

	p.Run()
}
