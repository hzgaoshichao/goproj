package main

import "fmt"

type StockObserver struct {
	Observer // 这里使用嵌套结构体的方式, 这样就可以提取出公共的方法和字段到嵌套结构体中了
}

func (so StockObserver) update() {
	fmt.Printf("%v: %v %v 请关闭股票行情,赶紧离开 \n", so.sub.getSubject().name, so.sub.getSubject().getAction(), so.Observer.name)
}

func NewStockObserver(name string, sub SubjectInterface) StockObserver {
	return StockObserver{
		Observer{name: name, sub: sub},
	}
}
