package main

type ConcreteIteratorDesc struct {
	aggregate *ConcreteAggregate
	current   int
}

func NewConcreteIteratorDesc(agg *ConcreteAggregate) *ConcreteIteratorDesc {

	return &ConcreteIteratorDesc{
		aggregate: agg,
		current:   agg.getCount() - 1,
	}
}

func (c *ConcreteIteratorDesc) first() any {
	return c.aggregate.getCurrentItem(c.aggregate.getCount() - 1)
}

func (c *ConcreteIteratorDesc) next() any {
	var ret any
	c.current--
	if c.current >= 0 {

		ret = c.aggregate.getCurrentItem(c.current)
	}
	return ret
}

func (c *ConcreteIteratorDesc) isDone() bool {
	if c.current <= 0 {
		return true
	} else {
		return false
	}
}

func (c *ConcreteIteratorDesc) currentItem() any {
	return c.aggregate.getCurrentItem(c.current)
}
