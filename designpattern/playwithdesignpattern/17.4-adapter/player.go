package main

type playerInterface interface {
	attach()
	defense()
}

type player struct {
	name string
}
