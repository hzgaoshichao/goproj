package main

type cashReturnRebateFactory struct {
	moneyRebate    float64
	moneyCondition int
	moneyReturn    int
}

func NewCashReturnRebateFactory(moneyCondition int, moneyReturn int, moneyRebate float64) *cashReturnRebateFactory {
	return &cashReturnRebateFactory{
		moneyRebate:    moneyRebate,
		moneyCondition: moneyCondition,
		moneyReturn:    moneyReturn,
	}
}

func (crr *cashReturnRebateFactory) createSalesMode() SaleInterface {
	//先满返, 再打折
	cn := NewCashNormal()
	cret := newCashReturn(crr.moneyCondition, crr.moneyReturn)
	creb := newCashRebate(crr.moneyRebate)

	creb.decorate(cn)
	cret.decorate(creb)

	return cret

}
