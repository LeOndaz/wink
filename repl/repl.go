package repl

import (
	"bufio"
	"fmt"
	"io"
	"lango/lexer"
	"lango/token"
)

const PROMPT = "$ "

func Start(in io.Reader, out io.Writer) {
	scanner := bufio.NewScanner(in)

	for {
		_, err := fmt.Fprintf(out, PROMPT)

		if err != nil {
			return
		}

		scanned := scanner.Scan()

		if scanned {
			line := scanner.Text()
			l := lexer.New(line)

			for t := l.NextToken(); t.Type != token.EOF; t = l.NextToken() {
				_, e := fmt.Fprintf(out, "%+v\n", t)
				if e != nil {
					return
				}
			}
		}
	}

}
