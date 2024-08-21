package main

import "fmt"

type origator struct {
	state string
}

func (o *origator) setState(value string) {
	o.state = value
}

func (o *origator) getState(value string) string {
	return o.state
}

func (o *origator) show() {
	fmt.Printf("State %v \n", o.state)
}

func (o *origator) createMemento() *memento {
	return &memento{state: o.state}
}

func (o *origator) recoverMemento(m *memento) {
	o.state = m.state
}
