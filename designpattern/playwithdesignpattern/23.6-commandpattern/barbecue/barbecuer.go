package barbecue

import "fmt"

type Barbecuer struct {
}

func (b *Barbecuer) bakeMutton() {
	fmt.Printf("烤羊肉串！\n")
}

func (b *Barbecuer) bakeChickenWing() {
	fmt.Printf("烤鸡翅！\n")
}
