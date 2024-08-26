package main

type AggregateInterface interface {
	CreateIterator() IteratorInterface
}
