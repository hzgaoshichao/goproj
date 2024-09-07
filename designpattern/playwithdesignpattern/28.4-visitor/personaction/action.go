package personaction

import (
	"fmt"
	"reflect"
)

type Action interface {
	GetManConclusion(concreteElementA *Man)
	GetWomanConclusion(concreteElementA *Woman)
}

type Success struct {
}

func (s *Success) GetManConclusion(concreteElementA *Man) {
	fmt.Printf("%v %v 时, 背后多半有一个伟大的女人. \n", reflect.TypeOf(concreteElementA), reflect.TypeOf(s))
}

func (s *Success) GetWomanConclusion(concreteElementA *Woman) {
	fmt.Printf("%v %v 时, 背后多半有一个不成功的男人. \n", reflect.TypeOf(concreteElementA), reflect.TypeOf(s))
}

type Failing struct {
}

func (f *Failing) GetManConclusion(concreteElementA *Man) {
	fmt.Printf("%v %v 时, 背后多半有一个伟大的女人. \n", reflect.TypeOf(concreteElementA), reflect.TypeOf(f))
}

func (f *Failing) GetWomanConclusion(concreteElementA *Woman) {
	fmt.Printf("%v %v 时, 背后多半有一个不成功的男人. \n", reflect.TypeOf(concreteElementA), reflect.TypeOf(f))
}

type Amativeness struct {
}

func (a *Amativeness) GetManConclusion(concreteElementA *Man) {
	fmt.Printf("%v %v 时, 背后多半有一个伟大的女人. \n", reflect.TypeOf(concreteElementA), reflect.TypeOf(a))
}

func (a *Amativeness) GetWomanConclusion(concreteElementA *Woman) {
	fmt.Printf("%v %v 时, 背后多半有一个不成功的男人. \n", reflect.TypeOf(concreteElementA), reflect.TypeOf(a))
}

type Marriage struct {
}

func (m *Marriage) GetManConclusion(concreteElementA *Man) {
	fmt.Printf("%v %v 时, 背后多半有一个伟大的女人. \n", reflect.TypeOf(concreteElementA), reflect.TypeOf(m))
}

func (m *Marriage) GetWomanConclusion(concreteElementA *Woman) {
	fmt.Printf("%v %v 时, 背后多半有一个不成功的男人. \n", reflect.TypeOf(concreteElementA), reflect.TypeOf(m))
}
