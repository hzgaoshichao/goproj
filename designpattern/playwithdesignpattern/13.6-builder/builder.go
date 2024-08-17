package main

type builderInterface interface {
	buildPartA()
	buildPartB()
	getResult() *product
}
