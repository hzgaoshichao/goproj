package main

import bb "bridge/bridgebasic"

func main() {
	var ab bb.AbstractionInterface
	ab = &bb.RefinedAbstraction{}
	ab.SetImplementor(&bb.ConcreteImplementorA{})
	ab.Operation()

	ab.SetImplementor(&bb.ConcreteImplementorB{})
	ab.Operation()

}
