package main

import (
	"fmt"
	"reflect"
)

type composite struct {
	component
	children []componentInterface
}

func NewComposite(name string) *composite {
	child := make([]componentInterface, 0)
	return &composite{
		component: component{name: name},
		children:  child,
	}
}

func (c *composite) add(value componentInterface) {
	c.children = append(c.children, value)
}

func (c *composite) remove(value componentInterface) {
	for i, child := range c.children {
		//if reflect.DeepEqual(value, child) // 这里涉及到结构体的比较, 需要特别注意 // 参考链接: https://segmentfault.com/a/1190000040099215
		if reflect.DeepEqual(value, child) {
			c.children = append(c.children[:i], c.children[i+1:]...)
		}
	}
}

func (c *composite) display(depth int) {
	for i := 0; i < depth; i++ {
		fmt.Printf("-")
	}

	fmt.Printf("%v \n", c.component.name)

	for _, item := range c.children {
		item.display(depth + 2)
	}

}
