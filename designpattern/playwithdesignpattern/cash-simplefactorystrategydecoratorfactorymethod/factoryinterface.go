package main

type factoryInterface interface {
	createSalesMode() SaleInterface
}
