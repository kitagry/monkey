package interp

import (
	"fmt"
	"io/ioutil"
	"os"

	"github.com/kitagry/monkey/evaluator"
	"github.com/kitagry/monkey/lexer"
	"github.com/kitagry/monkey/object"
	"github.com/kitagry/monkey/parser"
)

func Start(filename string) {
	f, err := os.Open(filename)
	if err != nil {
		fmt.Fprintf(os.Stderr, `"%s" doesn't found`, os.Args[1])
		os.Exit(1)
	}
	defer f.Close()

	data, err := ioutil.ReadAll(f)
	if err != nil {
		fmt.Fprintf(os.Stderr, "failed to read file: %v", err)
		os.Exit(1)
	}
	l := lexer.New(string(data))
	p := parser.New(l)
	program := p.ParseProgram()
	if len(p.Errors()) != 0 {
		for _, msg := range p.Errors() {
			fmt.Fprintf(os.Stderr, msg.Error())
		}
		os.Exit(1)
	}
	env := object.NewEnvironment()
	macroEnv := object.NewEnvironment()
	evaluator.DefineMacros(program, macroEnv)
	expanded := evaluator.ExpandMacros(program, macroEnv)

	evaluated := evaluator.Eval(expanded, env)
	if evaluated != nil {
		fmt.Println(evaluated.Inspect())
	}
}
