package main

type CashInterface interface {
	// 收取费用的抽象方法,参数我饿单价和数量
	acceptCash(price float64, num int) float64
}
