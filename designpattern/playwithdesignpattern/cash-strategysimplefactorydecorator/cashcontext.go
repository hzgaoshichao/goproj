package main

type CashContext struct {
	saleInterface SaleInterface
}

func NewCashContext(cashType int) CashContext {
	var saleInterface SaleInterface
	switch cashType {
	case 1:
		saleInterface = NewCashNormal()
	case 2:
		saleInterface = newCashRebate(0.8)
	case 3:
		saleInterface = newCashReturn(300, 100)
	case 4:
		//先打8折, 再满300返100
		cn := NewCashNormal()
		cret := newCashReturn(300, 100)
		creb := newCashRebate(0.8)

		cret.decorate(cn)
		creb.decorate(cret)

		saleInterface = creb
	}
	return CashContext{
		saleInterface: saleInterface,
	}
}

func (cc CashContext) GetResult(price float64, num int) float64 {
	return cc.saleInterface.acceptCash(price, num)
}
