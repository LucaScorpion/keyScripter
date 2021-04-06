package main

import (
	"fmt"
	"github.com/LucaScorpion/keyScripter/internal/parser"
	"github.com/alexflint/go-arg"
	"os"
)

var options struct {
	Script string `arg:"positional"`
	DryRun bool   `arg:"-d,--dry-run" help:"parse the script without running it"`
}

func main() {
	argParser := arg.MustParse(&options)

	if options.Script == "" {
		fmt.Printf("Error: no script file specified\n\n")
		argParser.WriteHelp(os.Stdout)
		os.Exit(1)
	}

	// Read the file.
	b, err := os.ReadFile(options.Script)
	if err != nil {
		if pathErr, ok := err.(*os.PathError); ok {
			fmt.Printf("An error occurred while trying to %s", pathErr.Error())
		} else {
			fmt.Printf("An error occurred while trying to read the script: %s", err)
		}
		os.Exit(1)
	}

	// Parse the script.
	script, err := parser.Parse(string(b))
	if err != nil {
		fmt.Printf("An error occurred while parsing the script: %s", err)
		os.Exit(1)
	}

	// Run the script.
	if !options.DryRun {
		script.Run()
	}
}
