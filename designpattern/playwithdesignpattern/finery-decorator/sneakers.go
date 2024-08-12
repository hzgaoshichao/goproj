package main

import "fmt"

type Sneakers struct {
	Finery
}

func (t *Sneakers) show() {
	fmt.Printf("Sneakers......")
	t.Finery.show()
}
