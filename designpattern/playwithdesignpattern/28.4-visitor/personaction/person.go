package personaction

import "reflect"

type Person interface {
	accept(visitor Action)
}

type Man struct {
}

func (m *Man) accept(visitor Action) {
	visitor.GetManConclusion(m) // 双分派技术
}

type Woman struct {
}

func (w *Woman) accept(visitor Action) {
	visitor.GetWomanConclusion(w) // 双分派技术
}

type ObjectStructure struct {
	elements []Person
}

func NewObjectStructure() *ObjectStructure {
	return &ObjectStructure{elements: make([]Person, 0)}
}

func (o *ObjectStructure) Attach(element Person) {
	o.elements = append(o.elements, element)
}

func (o *ObjectStructure) Detach(element Person) {
	for i, e := range o.elements {
		if reflect.DeepEqual(e, element) {
			o.elements = append(o.elements[:i], o.elements[i+1:]...)
		}
	}

}

func (o *ObjectStructure) Display(visitor Action) {
	for _, e := range o.elements {
		e.accept(visitor)
	}
}
