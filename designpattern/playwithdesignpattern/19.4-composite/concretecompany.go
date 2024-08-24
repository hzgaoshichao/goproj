package main

import (
	"fmt"
	"reflect"
)

type concreteCompany struct {
	name     string
	children []companyInterface
}

func NewConcreteCompany(name string) *concreteCompany {
	return &concreteCompany{
		name:     name,
		children: make([]companyInterface, 0),
	}
}

func (c *concreteCompany) add(com companyInterface) {
	c.children = append(c.children, com)
}

func (c *concreteCompany) remove(com companyInterface) {
	for i, item := range c.children {
		if reflect.DeepEqual(item, com) {
			c.children = append(c.children[:i], c.children[i+1:]...)
		}
	}
}

func (c *concreteCompany) display(depth int) {
	for i := 0; i < depth; i++ {
		fmt.Printf("-")
	}
	fmt.Printf("%v \n", c.name)

	for _, item := range c.children {
		item.display(depth + 2)
	}
}

func (c *concreteCompany) lineOfDuty() {

	for _, item := range c.children {
		item.lineOfDuty()
	}
}
