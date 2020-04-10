package repl

import (
	"bufio"
	"io"

	"github.com/kitagry/monkey/evaluator"
	"github.com/kitagry/monkey/lexer"
	"github.com/kitagry/monkey/parser"
)

const PROMPT = ">> "

func Start(in io.Reader, out io.Writer) {
	scanner := bufio.NewScanner(in)
	w := bufio.NewWriter(out)

	for {
		w.Write([]byte(PROMPT))
		w.Flush()
		scanned := scanner.Scan()
		if !scanned {
			return
		}

		line := scanner.Text()
		l := lexer.New(line)
		p := parser.New(l)

		program := p.ParseProgram()
		if len(p.Errors()) != 0 {
			printParserErrors(w, p.Errors())
			continue
		}

		evaluated := evaluator.Eval(program)
		if evaluated != nil {
			w.WriteString(evaluated.Inspect() + "\n")
			w.Flush()
		}
	}
}

func printParserErrors(w *bufio.Writer, errors []string) {
	for _, msg := range errors {
		w.WriteString(msg + "\n")
	}
	w.Flush()
}
