package repl

import (
	"bufio"
	"io"
	"strings"

	"github.com/kitagry/monkey/evaluator"
	"github.com/kitagry/monkey/lexer"
	"github.com/kitagry/monkey/object"
	"github.com/kitagry/monkey/parser"
)

const (
	PROMPT         = ">> "
	HALFWAY_PROMPT = ".. "
)

func Start(in io.Reader, out io.Writer, er io.Writer) {
	scanner := bufio.NewScanner(in)
	w := bufio.NewWriter(out)
	env := object.NewEnvironment()
	macroEnv := object.NewEnvironment()

	for {
		w.Write([]byte(PROMPT))
		w.Flush()

		var text string
		for {
			scanned := scanner.Scan()
			if !scanned {
				return
			}
			text += scanner.Text()
			text = strings.TrimSpace(text)
			if strings.Count(text, "(") == strings.Count(text, ")") &&
				strings.Count(text, "[") == strings.Count(text, "]") &&
				strings.Count(text, "{") == strings.Count(text, "}") &&
				strings.HasSuffix(text, ";") {
				break
			}
			w.Write([]byte(HALFWAY_PROMPT))
			w.Flush()
		}

		l := lexer.New(text)
		p := parser.New(l)
		program := p.ParseProgram()
		if len(p.Errors()) != 0 {
			printParserErrors(er, p.Errors())
			continue
		}

		evaluator.DefineMacros(program, macroEnv)
		expanded := evaluator.ExpandMacros(program, macroEnv)

		evaluated := evaluator.Eval(expanded, env)
		if evaluated != nil {
			w.WriteString(evaluated.Inspect() + "\n")
			w.Flush()
		}
	}
}

func printParserErrors(w io.Writer, errors []error) {
	for _, msg := range errors {
		w.Write([]byte(msg.Error() + "\n"))
	}
}
