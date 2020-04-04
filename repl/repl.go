package repl

import (
	"bufio"
	"fmt"
	"io"

	"github.com/kitagry/monkey/lexer"
	"github.com/kitagry/monkey/token"
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

		for tok := l.NextToken(); tok.Type != token.EOF; tok = l.NextToken() {
			w.Write([]byte(fmt.Sprintf("%+v\n", tok)))
		}
		w.Flush()
	}
}
