package main

import d "mediator/demo"

func main() {
	m := d.ConcreteMediator{}
	c1 := d.NewConcreteColleague1(&m)
	c2 := d.NewConcreteColleague2(&m)

	m.SetColleague1(c1)
	m.SetColleague2(c2)

	c1.Send("吃过饭了吗?")
	c2.Send("没有呢, 你打算请客?")
}
