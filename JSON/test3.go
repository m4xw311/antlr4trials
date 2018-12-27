
package main

import (
	"github.com/antlr/antlr4/runtime/Go/antlr"
	"./parser"
	"os"
)

type jsonListener struct {
	*parser.BaseJSONListener
}

func main() {
	input, _ := antlr.NewFileStream(os.Args[1])
	lexer := parser.NewJSONLexer(input)
	stream := antlr.NewCommonTokenStream(lexer, antlr.TokenDefaultChannel)
	p := parser.NewJSONParser(stream)
	antlr.ParseTreeWalkerDefault.Walk(&jsonListener{}, p.Json())
}
