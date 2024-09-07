package basic

import (
	"fmt"
	"reflect"
)

type Visitor interface {
	VisitConcreteElementA(e *ConcreteElementA)
	VisitConcreteElementB(e *ConcreteElementB)
}

type ConcreteVisitor1 struct {
}

func (v *ConcreteVisitor1) VisitConcreteElementA(e *ConcreteElementA) {
	fmt.Printf("%v 被 %v访问 \n", reflect.TypeOf(e), reflect.TypeOf(v))
}

func (v *ConcreteVisitor1) VisitConcreteElementB(e *ConcreteElementB) {
	fmt.Printf("%v 被 %v访问 \n", reflect.TypeOf(e), reflect.TypeOf(v))
}

type ConcreteVisitor2 struct {
}

func (v *ConcreteVisitor2) VisitConcreteElementA(e *ConcreteElementA) {
	fmt.Printf("%v 被 %v访问 \n", reflect.TypeOf(e), reflect.TypeOf(v))
}

func (v *ConcreteVisitor2) VisitConcreteElementB(e *ConcreteElementB) {
	fmt.Printf("%v 被 %v访问 \n", reflect.TypeOf(e), reflect.TypeOf(v))
}
