package main

import "fmt"

type Person struct {
	name string
}

func NewPerson(name string) Person {
	return Person{name: name}
}

func (p *Person) show() {
	fmt.Printf("装扮的: %s", p.name)
}
