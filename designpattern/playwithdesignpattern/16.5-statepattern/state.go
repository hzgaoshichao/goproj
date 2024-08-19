package main

type StateInterface interface {
	handle(context *Context)
}

type ConcreteStateA struct {
}

func (c *ConcreteStateA) handle(context *Context) {
	context.setState(&ConcreteStateB{})
}

type ConcreteStateB struct {
}

func (c *ConcreteStateB) handle(context *Context) {
	context.setState(&ConcreteStateA{})
}
