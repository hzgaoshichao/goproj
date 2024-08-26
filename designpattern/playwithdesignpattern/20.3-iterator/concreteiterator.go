package main

type ConcreteIterator struct {
	aggregate *ConcreteAggregate
	current   int
}

func NewConcreteIterator(agg *ConcreteAggregate) *ConcreteIterator {
	return &ConcreteIterator{
		aggregate: agg,
		current:   0,
	}
}

func (c *ConcreteIterator) first() any {
	return c.aggregate.getCurrentItem(0)
}

func (c *ConcreteIterator) next() any {
	var ret any
	c.current++
	if c.current < c.aggregate.getCount() {

		ret = c.aggregate.getCurrentItem(c.current)
	}
	return ret
}

func (c *ConcreteIterator) isDone() bool {
	if c.current >= c.aggregate.getCount() {
		return true
	} else {
		return false

	}
}

func (c *ConcreteIterator) currentItem() any {
	return c.aggregate.getCurrentItem(c.current)
}
