package main

import "fmt"

type adaptee struct {
}

func (a *adaptee) specificRequest() {
	fmt.Printf("特殊请求!")
}
