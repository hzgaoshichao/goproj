package main

import b "visitor/basic"

func main() {
	o := b.NewObjectStructure()
	o.Attach(&b.ConcreteElementA{})
	o.Attach(&b.ConcreteElementB{})

	v1 := b.ConcreteVisitor1{}
	v2 := b.ConcreteVisitor2{}
	o.Accept(&v1)
	o.Accept(&v2)
}
