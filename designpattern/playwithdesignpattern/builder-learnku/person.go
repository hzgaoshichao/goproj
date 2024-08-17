package main

import "fmt"

type PersonInterface interface {
	Speak(message string)
	Sleep()
}

type Person struct {
	name    string
	age     int
	gender  string
	address string
}

func (p *Person) Speak(msg string) {
	fmt.Printf("%s says: %s\n", p.name, msg)
}

func (p *Person) Sleep() {
	fmt.Printf("%s is sleeping now...\n", p.name)
}
