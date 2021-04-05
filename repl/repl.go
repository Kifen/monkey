package repl

import (
	"bufio"
	"fmt"
	"io"
)

const PROMPT = ">> "

func start(in io.Reader, out io.Writer) {
	scanner := bufio.NewScanner(in)

	for {
		fmt.Printf(PROMPT)
		scanned := scanner.Scan()
	}
}
