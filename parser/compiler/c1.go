package main

import (
	"fmt"
	"os"

	"github.com/antlr4-go/antlr/v4"
	"github.com/xt0fer/smog5/parser"
)

type somListener struct {
	*parser.BaseSOMListener
}

func main() {
	// Setup the input
	//is := antlr.NewInputStream("Hello = ( run = ('Hello, world from smog5!' println ) )")
	if len(os.Args) < 2 {
		fmt.Println("Missing parameter, provide file name!")
		return
	}
	is, err := antlr.NewFileStream(os.Args[1])
	if err != nil {
		fmt.Println("Can't read file:", os.Args[1])
		panic(err)
	}

	// Create the Lexer
	lexer := parser.NewSOMLexer(is)
	stream := antlr.NewCommonTokenStream(lexer, antlr.TokenDefaultChannel)

	// Create the Parser
	p := parser.NewSOMParser(stream)

	// Finally parse the expression
	antlr.ParseTreeWalkerDefault.Walk(&somListener{}, p.Classdef())

}
