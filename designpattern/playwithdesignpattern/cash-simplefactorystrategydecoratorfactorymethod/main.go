package main

import "fmt"

func main() {
	var price float64 = 500
	var num int = 4
	var discount int = 5
	cashContext := NewCashContext(discount)
	totalPrices := cashContext.GetResult(price, num)
	fmt.Printf("totalPrices: %f", totalPrices)
}
