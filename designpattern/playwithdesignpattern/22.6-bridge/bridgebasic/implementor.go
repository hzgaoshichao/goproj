package bridgebasic

import "fmt"

type ImplementorInterface interface {
	Operation()
}

type ConcreteImplementorA struct {
}

func (ci *ConcreteImplementorA) Operation() {
	fmt.Printf("具体实现 A 的方法执行 \n")
}

type ConcreteImplementorB struct {
}

func (ci *ConcreteImplementorB) Operation() {
	fmt.Printf("具体实现 B 的方法执行 \n")
}
