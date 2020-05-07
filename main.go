package main

import (
	"flag"
	"fmt"
	"os"

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

		n := nullWriter{}
		repl.Start(file, &n, os.Stderr)
	}
}

type nullWriter struct {
}

func (nw *nullWriter) Write(p []byte) (n int, err error) {
	return
}
