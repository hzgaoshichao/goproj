package main

import (
	"fmt"
	"reflect"
)

type Context struct {
	state StateInterface
}

func NewContext(state StateInterface) Context {
	return Context{
		state: state,
	}
}

func (c *Context) getState() StateInterface {
	return c.state
}

func (c *Context) setState(s StateInterface) {
	c.state = s
	fmt.Printf("当前状态: %v \n", reflect.TypeOf(c.state))
}

func (c *Context) request() {
	c.state.handle(c)
}
