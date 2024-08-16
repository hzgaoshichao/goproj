package main

type Facade struct {
	one   *subSystemOne
	two   *subSystemTwo
	three *subSystemThree
	four  *SubSystemFour
}

func NewFacade() *Facade {
	one := &subSystemOne{}
	two := &subSystemTwo{}
	three := &subSystemThree{}
	four := &SubSystemFour{}

	return &Facade{
		one:   one,
		two:   two,
		three: three,
		four:  four,
	}
}

func (f *Facade) methodA() {
	f.one.methodOne()
	f.two.methodTwo()
	f.three.methodThree()
	f.four.methodFour()
}

func (f *Facade) methodB() {
	f.two.methodTwo()
	f.three.methodThree()

}
