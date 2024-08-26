package main

import "fmt"

func main() {
	bus := NewConcreteAggregate()
	bus.add("大鸟")
	bus.add("小菜")
	bus.add("行李")
	bus.add("老外")
	bus.add("公交内部员工")
	bus.add("小偷")

	// 正序迭代器
	// conductor := NewConcreteIterator(bus)
	// conductor := bus.createIterator()

	// 倒序迭代器
	conductor := NewConcreteIteratorDesc(bus)
	conductor.first()
	for {
		if conductor.isDone() {
			break
		}
		fmt.Printf("%v ,请买票上车 \n", conductor.currentItem())
		conductor.next()
	}
}
