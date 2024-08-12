package main

import "fmt"

func main() {
	var price float64 = 500
	var num int = 50
	var discount int = 3
	cash := createCashAccept(discount)
	totalPrices := cash.acceptCash(price, num)
	fmt.Printf("totalPrices: %f", totalPrices)
}
