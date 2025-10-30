package repl

import (
	"MonkeyInterpreter/token"
	"MonkeyInterpreter/tokenizer"
	"bufio"
	"fmt"
	"io"
)

const PROMPT = ">> "

func Start(in io.Reader) {
	scanner := bufio.NewScanner(in)

	for {
		fmt.Printf(PROMPT)
		input := scanner.Scan()
		if !input {
			return
		}

		line := scanner.Text()
		t := tokenizer.NewTokenizer(line)
		for tok := t.NextToken(); tok.Type != token.EOF; tok = t.NextToken() {
			fmt.Printf("%#v\n", tok)
		}
	}
}
