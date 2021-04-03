package main

import (
	"github.com/LucaScorpion/tas-scripter/internal/parser"
	"io/ioutil"
)

func main() {
	b, _ := ioutil.ReadFile("example.txt")
	parser.Parse(string(b))
}
