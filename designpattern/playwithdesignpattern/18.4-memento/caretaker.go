package main

type caretaker struct {
	memento *memento
}

func (c *caretaker) getMemento() *memento {
	return c.memento
}

func (c *caretaker) setMemento(value *memento) {
	c.memento = value
}
