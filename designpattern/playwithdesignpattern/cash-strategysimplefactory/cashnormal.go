package main

type CashNormal struct {
}

func newCashNormal() CashNormal {
	return CashNormal{}
}
func (cn CashNormal) acceptCash(price float64, num int) float64 {
	return price * float64(num)
}
