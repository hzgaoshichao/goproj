package basic

import "reflect"

type Element interface {
	accept(vistor Visitor)
}

type ConcreteElementA struct {
}

func (e *ConcreteElementA) accept(vistor Visitor) {
	vistor.VisitConcreteElementA(e)
}

func (e *ConcreteElementA) OperationB() {

}

type ConcreteElementB struct {
}

func (e *ConcreteElementB) accept(vistor Visitor) {
	vistor.VisitConcreteElementB(e)
}

func (e *ConcreteElementB) OperationB() {

}

type ObjectStructure struct {
	elements []Element
}

func NewObjectStructure() *ObjectStructure {
	return &ObjectStructure{
		elements: make([]Element, 0),
	}
}

func (o *ObjectStructure) Attach(e Element) {
	o.elements = append(o.elements, e)
}

func (o *ObjectStructure) Detach(element Element) {
	for i, e := range o.elements {
		if reflect.DeepEqual(e, element) {
			o.elements = append(o.elements[:i], o.elements[i+1:]...)
		}
	}
}

func (o *ObjectStructure) Accept(Visitor Visitor) {
	for _, e := range o.elements {
		e.accept(Visitor)
	}

}
