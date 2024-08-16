package main

type Facade struct {
	one   *subSystemOne
	two   *subSystemTwo
	three *subSystemThree
}

func NewFacade() *Facade {
	one := &subSystemOne{}
	two := &subSystemTwo{}
	three := &subSystemThree{}

	return &Facade{
		one:   one,
		two:   two,
		three: three,
	}
}

func (f *Facade) methodA() {
	f.one.methodOne()
	f.two.methodTwo()
	f.three.methodThree()
}

func (f *Facade) methodB() {
	f.two.methodTwo()
	f.three.methodThree()
}
