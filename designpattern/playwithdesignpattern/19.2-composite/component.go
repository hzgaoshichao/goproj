package main

type componentInterface interface {
	add(comp componentInterface)
	remove(com componentInterface)
	display(depth int)
}

type component struct {
	name string
}
