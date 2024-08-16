package main

import "fmt"

type subSystemOne struct {
}

func (s subSystemOne) methodOne() {
	fmt.Printf("子系统方法一")
}
