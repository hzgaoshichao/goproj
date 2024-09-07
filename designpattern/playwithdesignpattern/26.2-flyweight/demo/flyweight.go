package demo

import "fmt"

type Flyweight interface {
	Operation(extrinsicstate int)
}

type ConcreteFlyweight struct {
}

func (cf *ConcreteFlyweight) Operation(extrinsicstate int) {
	fmt.Printf("具体Flyweight: %v \n", extrinsicstate)
}

type UnsharedConcreteFlyweight struct {
}

func (cf *UnsharedConcreteFlyweight) Operation(extrinsicstate int) {
	fmt.Printf("不共享的具体Flyweight: %v \n", extrinsicstate)
}

type FlyweightFactory struct {
	flyweights map[string]*ConcreteFlyweight
}

func NewFlyweightFactory() FlyweightFactory {
	flyweightFactory := FlyweightFactory{}
	flyweightFactory.flyweights = make(map[string]*ConcreteFlyweight)
	flyweightFactory.flyweights["X"] = &ConcreteFlyweight{}
	flyweightFactory.flyweights["Y"] = &ConcreteFlyweight{}
	flyweightFactory.flyweights["Z"] = &ConcreteFlyweight{}
	return flyweightFactory
}

func (f *FlyweightFactory) GetFlyweight(key string) Flyweight {
	return f.flyweights[key]
}
