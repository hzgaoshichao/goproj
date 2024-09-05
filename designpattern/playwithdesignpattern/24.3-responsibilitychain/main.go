package main

import rc "responsibilitychain/respchain"

func main() {
	h1 := rc.ConcreteHandler1{}
	h2 := rc.ConcreteHandler2{}
	h3 := rc.ConcreteHandler3{}

	h1.SetSuccessor(&h2)
	h2.SetSuccessor(&h3)
	request := []int{2, 5, 14, 22, 18, 3, 27, 20}
	for _, request := range request {
		h1.HandleRequest(request)
	}
}
