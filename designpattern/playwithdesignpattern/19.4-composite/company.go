package main

type companyInterface interface {
	add(com companyInterface)
	remove(com companyInterface)
	display(depth int)
	lineOfDuty()
}
