package main

type CashRebate struct {
	moneyRebate float64
}

func newCashRebate(moneyRebate float64) CashRebate {
	return CashRebate{
		moneyRebate: moneyRebate,
	}
}

func (cr CashRebate) acceptCash(price float64, num int) float64 {
	return price * float64(num) * cr.moneyRebate
}
