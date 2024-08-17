package main

import "fmt"

type product struct {
	parts []string
}

func (p *product) add(part string) {
	p.parts = append(p.parts, part)
}

func (p *product) show() {
	for _, part := range p.parts {
		fmt.Printf("%v \n", part)
	}
}
