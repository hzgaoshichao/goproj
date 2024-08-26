package main

type IteratorInterface interface {
	first() any
	next() any
	isDone() bool
	currentItem() any
}
