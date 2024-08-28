package bridgebasic

import "fmt"

type AbstractionInterface interface {
	SetImplementor(implementor ImplementorInterface)
	Operation()
}

type RefinedAbstraction struct {
	implementor ImplementorInterface
}

func (ra *RefinedAbstraction) Operation() {
	fmt.Printf(" 具体的 Abstraction \n")
	ra.implementor.Operation()
}

func (ra *RefinedAbstraction) SetImplementor(implementor ImplementorInterface) {
	ra.implementor = implementor
}
