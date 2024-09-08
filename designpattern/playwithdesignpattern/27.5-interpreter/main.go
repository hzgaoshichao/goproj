package main

import (
	"fmt"
	m "interpreter/musicinterpreter"
)

func main() {
	context := &m.PlayContext{}
	fmt.Printf("音乐-上海滩 \n")

	// context.SetPlayText("O 2 E 0.5 G 0.5 A 3 E 0.5 G 0.5 D 3 E 0.5 G 0.5 A 0.5 O 3 C 1 O 2 A 0.5 G 1 C 0.5 E 0.5 D 3 ")

	context.SetPlayText("T 500 O 2 E 0.5 G 0.5 A 3 E 0.5 G 0.5 D 3 E 0.5 G 0.5 A 0.5 O 3 C 1 O 2 A 0.5 G 1 C 0.5 E 0.5 D 3 ")

	var expression m.Express

	for len(context.GetPlayText()) > 0 {
		str := context.GetPlayText()[0:1]
		switch str {
		case "O":
			expression = &m.Scale{}
		case "C":
			fallthrough
		case "D":
			fallthrough
		case "E":
			fallthrough
		case "F":
			fallthrough
		case "G":
			fallthrough
		case "A":
			fallthrough
		case "B":
			fallthrough
		case "P":
			expression = &m.Note{}
		case "T":
			expression = &m.Speed{}

		}
		expression.Interpret(context, expression)
	}
}
