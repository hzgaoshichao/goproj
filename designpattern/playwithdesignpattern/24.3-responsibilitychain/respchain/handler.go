package respchain

import (
	"fmt"
	"reflect"
)

type HandlerInterface interface {
	SetSuccessor(successor HandlerInterface)
	HandleRequest(request int)
}

type ConcreteHandler struct {
	successor HandlerInterface
}

func (ch *ConcreteHandler) SetSuccessor(successor HandlerInterface) {
	ch.successor = successor
}

func (ch *ConcreteHandler) HandleRequest(request int) {
	fmt.Println("Not Implemented")
}

type ConcreteHandler1 struct {
	ConcreteHandler
}

func (ch *ConcreteHandler1) HandleRequest(request int) {
	if request >= 0 && request < 10 {
		fmt.Printf("%v 处理请求 %v \n", reflect.TypeOf(ch), request)
	} else if ch.successor != nil {
		ch.successor.HandleRequest(request)
	}
}

type ConcreteHandler2 struct {
	ConcreteHandler
}

func (ch *ConcreteHandler2) HandleRequest(request int) {
	if request >= 10 && request < 20 {
		fmt.Printf("%v 处理请求 %v \n", reflect.TypeOf(ch), request)
	} else if ch.successor != nil {
		ch.successor.HandleRequest(request)
	}
}

type ConcreteHandler3 struct {
	ConcreteHandler
}

func (ch *ConcreteHandler3) HandleRequest(request int) {
	if request >= 20 && request < 30 {
		fmt.Printf("%v 处理请求 %v \n", reflect.TypeOf(ch), request)
	} else if ch.successor != nil {
		ch.successor.HandleRequest(request)
	}
}
