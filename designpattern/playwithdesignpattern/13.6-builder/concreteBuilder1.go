package main

type concreteBuilder1 struct {
	product *product
}

func (c *concreteBuilder1) buildPartA() {
	c.product.add("部件 A")
}

func (c *concreteBuilder1) buildPartB() {
	c.product.add("部件 B")
}

func (c *concreteBuilder1) getResult() *product {
	return c.product
}
