package main

type CashSuper struct {
	saleInterface SaleInterface
}

func (cs *CashSuper) decorate(saleInterface SaleInterface) {
	cs.saleInterface = saleInterface

}

func (cs *CashSuper) acceptCash(price float64, num int) float64 {
	var result float64
	if cs.saleInterface != nil {
		result = cs.saleInterface.acceptCash(price, num)
	}
	return result
}
