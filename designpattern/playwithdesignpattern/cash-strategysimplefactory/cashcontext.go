package main

type CashContext struct {
	cash CashInterface
}

func NewCashContext(cashType int) CashContext {
	var cash CashInterface
	switch cashType {
	case 1:
		cash = newCashNormal()
	case 2:
		cash = newCashRebate(0.8)
	case 3:
		cash = newCashReturn(300, 100)
	}
	return CashContext{
		cash: cash,
	}
}

func (cc CashContext) GetResult(price float64, num int) float64 {
	return cc.cash.acceptCash(price, num)
}
