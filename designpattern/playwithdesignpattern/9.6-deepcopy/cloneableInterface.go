package main

type cloneableInterface interface {
	clone() cloneableInterface
}
