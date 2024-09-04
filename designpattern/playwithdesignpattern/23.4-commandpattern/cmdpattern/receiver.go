package cmdpattern

import "fmt"

type Receiver struct {
}

func (r *Receiver) action() {
	fmt.Printf("执行请求！")
}
