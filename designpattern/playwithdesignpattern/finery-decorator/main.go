package main

import "fmt"

func main() {
	xc := NewPerson("小菜")

	fmt.Printf("第一种装扮......")
	pqx := new(Sneakers)
	pqx.decorate(&xc)

	dtx := Tshirts{}
	dtx.decorate(pqx)

	dtx.show()

}
