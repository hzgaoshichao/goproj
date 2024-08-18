package main

import "fmt"

type NBAObserver struct {
	Observer
}

func (n NBAObserver) update() {
	fmt.Printf("%v: %v %v 请关闭股票行情,赶紧离开 \n", n.sub.getSubject().name, n.sub.getSubject().getAction(), n.Observer.name)
}

func NewNBAObserver(name string, sub SubjectInterface) NBAObserver {
	return NBAObserver{
		Observer{name: name, sub: sub},
	}
}
