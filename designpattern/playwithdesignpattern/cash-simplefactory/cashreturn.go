package main

import (
	"math"
)

type CashReturn struct {
	//返利条件, 比如满300,减去100
	moneyCondition int
	moneyReturn    int
}

func newCashReturn(moneyCondition int, moneyReturn int) CashReturn {
	return CashReturn{
		moneyCondition: moneyCondition,
		moneyReturn:    moneyReturn,
	}
}

func (cr CashReturn) acceptCash(price float64, num int) float64 {
	result := price * float64(num)
	if cr.moneyCondition > 0 && result >= float64(cr.moneyCondition) {
		result = result - math.Floor(float64(result)/float64(cr.moneyCondition))*float64(cr.moneyReturn)
	}
	return result
}
