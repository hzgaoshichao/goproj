package main

import (
	"fmt"
	c "simplefactory/calculator"
)

func main() {

	var a, b float64
	fmt.Println("请输入两个数：")
	fmt.Scanf("%f %f", &a, &b)
	fmt.Println("请输入运算符号(+、-、*、/):")
	var operate string
	fmt.Scanf("%s", &operate)

	oper := c.CreateOperate(operate)
	result := oper.GetResult(a, b)

	fmt.Printf("%f %s %f = %f\n", a, operate, b, result)

}
