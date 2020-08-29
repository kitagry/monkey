package main

import (
	"flag"
	"fmt"
	"io/ioutil"
	"os"

	"github.com/kitagry/monkey/evaluator"
	"github.com/kitagry/monkey/lexer"
	"github.com/kitagry/monkey/object"
	"github.com/kitagry/monkey/parser"
	"github.com/kitagry/monkey/repl"
)

func main() {
	flag.Parse()
	args := flag.Args()

	switch len(args) {
	case 0:
		repl.Start(os.Stdin, os.Stdout, os.Stderr)
	case 1:
		file, err := os.Open(os.Args[1])
		if err != nil {
			fmt.Fprintf(os.Stderr, `"%s" doesn't found`, os.Args[1])
			os.Exit(1)
		}
		defer file.Close()

		data, err := ioutil.ReadAll(file)
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
}

type nullWriter struct {
}

func (nw *nullWriter) Write(p []byte) (n int, err error) {
	return
}
