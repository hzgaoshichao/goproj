package main

import "fmt"

type guards struct {
	player
}

func NewGuards(name string) *guards {
	return &guards{
		player{
			name: name,
		},
	}
}

func (g *guards) attack() {
	fmt.Printf("guards attack!!!! \n")
}

func (g *guards) defense() {
	fmt.Printf("guards defense!!!! \n")
}
