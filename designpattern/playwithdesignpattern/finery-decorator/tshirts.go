package main

import "fmt"

type Tshirts struct {
	Finery
}

func (t *Tshirts) show() {
	fmt.Printf("大 T Shirt")
	t.Finery.show()
}
