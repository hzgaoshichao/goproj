package main

type CashContext struct {
	saleInterface SaleInterface
}

func NewCashContext(cashType int) CashContext {
	var fs factoryInterface
	switch cashType {
	case 1: // 原价
		fs = NewCashRebateReturnFactory(1, 0, 0)
	case 2: // 打八折
		fs = NewCashRebateReturnFactory(0.8, 0, 0)
	case 3: // 打七折
		fs = NewCashRebateReturnFactory(0.7, 0, 0)
	case 4: // 满300减100
		fs = NewCashRebateReturnFactory(1, 300, 100)
	case 5: // 先打八折, 再满300减100
		fs = NewCashRebateReturnFactory(0.8, 300, 100)
	case 6: // 再满200减50, 先打七折
		fs = NewCashReturnRebateFactory(200, 50, 0.7)
	}

	return CashContext{
		saleInterface: fs.createSalesMode(),
	}
}

func (cc CashContext) GetResult(price float64, num int) float64 {
	return cc.saleInterface.acceptCash(price, num)
}
