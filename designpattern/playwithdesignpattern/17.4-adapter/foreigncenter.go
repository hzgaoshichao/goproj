package main

import "fmt"

type foreignCenter struct {
	fname string
}

func NewForeignCenter(name string) *foreignCenter {
	return &foreignCenter{
		fname: name,
	}
}

func (f *foreignCenter) foreignCenterAttack() {
	fmt.Printf("外籍中锋 %v 进攻 \n", f.fname)
}

func (f *foreignCenter) foreignCenterDefense() {
	fmt.Printf("外籍中锋 %v 防御 \n", f.fname)
}
