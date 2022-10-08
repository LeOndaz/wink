package main

import (
	"lango/repl"
	"os"
)

func main() {
	repl.Start(os.Stdin, os.Stdout)
}
