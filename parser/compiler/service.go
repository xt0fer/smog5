package main

import (
	"fmt"

	"github.com/antlr4-go/antlr/v4"
)

func (s *somListener) VisitTerminal(node antlr.TerminalNode) {
	fmt.Printf("%v\n", node.GetText(), node.GetParent())
}

