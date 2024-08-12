package main

type SaleInterface interface {
	acceptCash(price float64, num int) float64
}
