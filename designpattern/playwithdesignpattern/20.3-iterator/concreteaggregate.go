package main

type ConcreteAggregate struct {
	items []any
}

func NewConcreteAggregate() *ConcreteAggregate {
	return &ConcreteAggregate{
		items: make([]any, 0),
	}
}

func (c *ConcreteAggregate) createIterator() IteratorInterface {
	return NewConcreteIterator(c)
}

func (c *ConcreteAggregate) getCount() int {
	return len(c.items)
}

func (c *ConcreteAggregate) add(object any) {
	c.items = append(c.items, object)
}

func (c *ConcreteAggregate) getCurrentItem(index int) any {
	return c.items[index]
}
