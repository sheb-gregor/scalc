package main

import (
	"os"
)

func main() {
	state := LexerCtx{}
	state.Parse(os.Args[1:])

	result := state.root.Eval()
	for _, val := range result {
		println(val)
	}
}
