package main

type CashRebate struct {
	CashSuper
	moneyRebate float64
}

func newCashRebate(moneyRebate float64) *CashRebate {
	return &CashRebate{
		moneyRebate: moneyRebate,
	}
}

func (cr *CashRebate) acceptCash(price float64, num int) float64 {
	result := price * float64(num) * cr.moneyRebate
	return cr.CashSuper.acceptCash(result, 1)
}
