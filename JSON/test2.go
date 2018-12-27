
package main

import (
	"github.com/antlr/antlr4/runtime/Go/antlr"
	"./parser"
	"os"
	"fmt"
)

func main() {
	input, _ := antlr.NewFileStream(os.Args[1])
	lexer := parser.NewJSONLexer(input)
	for {
		t:=lexer.NextToken()
		if t.GetTokenType() == antlr.TokenEOF {
			break
		}
		fmt.Printf("%s (%q)\n",
			lexer.SymbolicNames[t.GetTokenType()], t.GetText())
	}
}
