package main

import (
	"fmt"
	"github.com/LucaScorpion/keyScripter/internal/parser"
	"io/ioutil"
	"os"
)

func main() {
	b, _ := ioutil.ReadFile("simple.txt")

	script, err := parser.Parse(string(b))
	if err != nil {
		fmt.Printf("An error occurred while parsing the script: %s", err)
		os.Exit(1)
	}

	script.Run()
}
