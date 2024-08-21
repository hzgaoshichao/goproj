package main

import "fmt"

type forward struct {
	player
}

func NewForward(name string) forward {
	return forward{
		player{
			name: name,
		},
	}
}

func (f *forward) attack() {
	fmt.Printf("Forward attack!!!! \n")
}

func (f *forward) defense() {
	fmt.Printf("Forward defense!!!! \n")
}
