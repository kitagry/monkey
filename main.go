package main

import (
	"flag"
	"os"

	"github.com/kitagry/monkey/interp"
	"github.com/kitagry/monkey/repl"
)

func main() {
	flag.Parse()
	args := flag.Args()

	switch len(args) {
	case 0:
		repl.Start(os.Stdin, os.Stdout, os.Stderr)
	case 1:
		interp.Start(os.Args[1])
	}
}
