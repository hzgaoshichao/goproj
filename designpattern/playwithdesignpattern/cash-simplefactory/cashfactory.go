package main

type CashFactory struct {
}

func createCashAccept(cashType int) CashInterface {
	var cash CashInterface
	switch cashType {
	case 1:
		cash = newCashNormal()
	case 2:
		cash = newCashRebate(0.8)
	case 3:
		cash = newCashReturn(300, 100)
	}
	return cash
}
