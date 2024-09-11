package calculator

import "fmt"

type Operation interface {
	GetResult(numberA float64, numberB float64) float64
}

type Add struct {
}

func (a *Add) GetResult(numberA float64, numberB float64) float64 {
	return numberA + numberB

}

type Sub struct {
}

func (a *Sub) GetResult(numberA float64, numberB float64) float64 {
	return numberA - numberB

}

type Mul struct {
}

func (a *Mul) GetResult(numberA float64, numberB float64) float64 {
	return numberA * numberB

}

type Div struct {
}

func (a *Div) GetResult(numberA float64, numberB float64) float64 {
	if numberB == 0 {
		fmt.Println("除数不能为0")
		panic("除数不能为0")
	}
	return numberA / numberB

}

func CreateOperate(operate string) Operation {
	var oper Operation
	switch operate {
	case "+":
		oper = &Add{}
	case "-":
		oper = &Sub{}
	case "*":
		oper = &Mul{}
	case "/":
		oper = &Div{}

	}
	return oper
}
