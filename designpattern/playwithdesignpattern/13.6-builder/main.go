package main

func main() {
	director := director{}
	b1 := concreteBuilder1{product: &product{parts: make([]string, 0)}}
	director.construct(&b1)
	p1 := b1.getResult()
	p1.show()

	b2 := concreteBuilder2{product: &product{parts: make([]string, 0)}}
	director.construct(&b2)
	p2 := b2.getResult()
	p2.show()
}
