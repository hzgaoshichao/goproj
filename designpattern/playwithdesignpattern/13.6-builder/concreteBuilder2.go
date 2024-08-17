package main

type concreteBuilder2 struct {
	product *product
}

func (c *concreteBuilder2) buildPartA() {
	c.product.add("部件 X")
}

func (c *concreteBuilder2) buildPartB() {
	c.product.add("部件 Y")
}

func (c *concreteBuilder2) getResult() *product {
	return c.product
}
