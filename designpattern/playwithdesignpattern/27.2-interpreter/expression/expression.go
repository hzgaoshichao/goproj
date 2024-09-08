package expression

import "fmt"

type Context struct {
	input  string
	output string
}

func (c *Context) GetInput() string {
	return c.input
}

func (c *Context) SetInput(ipt string) {
	c.input = ipt
}

func (c *Context) GetOutput() string {
	return c.output
}

func (c *Context) SetOutput(opt string) {
	c.output = opt
}

type Expression interface {
	Interpret(ctx *Context)
}

type TerminalExpression struct {
}

func (t *TerminalExpression) Interpret(ctx *Context) {
	fmt.Printf("终端解释器 \n")
}

type NonterminalExpression struct {
}

func (nt *NonterminalExpression) Interpret(ctx *Context) {
	fmt.Printf("非终端解释器 \n")
}
