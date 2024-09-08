package main

import ie "interpreter/expression"

func main() {
	context := ie.Context{}
	list := make([]ie.Expression, 0)
	list = append(list, &ie.TerminalExpression{})
	list = append(list, &ie.NonterminalExpression{})
	list = append(list, &ie.TerminalExpression{})
	list = append(list, &ie.TerminalExpression{})

	for _, exp := range list {
		exp.Interpret(&context)
	}
}
