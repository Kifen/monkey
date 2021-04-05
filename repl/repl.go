package repl

import (
	"bufio"
	"fmt"
	"io"
	"log"
	"os"

	"github.com/Kifen/monkey/lexer"
	"github.com/Kifen/monkey/token"
)

const PROMPT = ">> "

func Start(in io.Reader, out io.Writer) {

	scanner := bufio.NewScanner(in)

	for {
		fmt.Printf(PROMPT)
		scanned := scanner.Scan()
		if !scanned {
			return
		}

		fileName := scanner.Text()

		file, err := os.Open(fileName)
		if err != nil {
			log.Panicf("Failed to open file %s", err)
		}
		l := lexer.New(file)

		for tok := l.NextToken(); tok.Type != token.EOF; tok = l.NextToken() {
			fmt.Printf("%+v\n", tok)
		}
	}
}
