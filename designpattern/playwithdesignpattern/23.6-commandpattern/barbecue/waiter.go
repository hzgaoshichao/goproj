package barbecue

import (
	"fmt"
	"reflect"
	"time"
)

type Waiter struct {
	orders []CommandInterface
}

func NewWaiter() *Waiter {
	return &Waiter{orders: make([]CommandInterface, 0)}
}

func (w *Waiter) SetOrder(cmd CommandInterface) {
	structName := reflect.TypeOf(cmd).String()
	if structName == "*barbecue.BakeChickenWingCommand" {
		fmt.Printf("服务员：鸡翅没有了，请点别的烧烤。\n")
	} else {
		w.orders = append(w.orders, cmd)
		fmt.Printf("增加订单：%v 时间：%v \n", structName, time.Now().String())
	}
}

func (w *Waiter) CancelOrder(cmd CommandInterface) {
	structName := reflect.TypeOf(cmd).String()
	for i, c := range w.orders {
		if reflect.DeepEqual(cmd, c) {
			w.orders = append(w.orders[:i], w.orders[i+1:]...)
			break
		}
	}
	fmt.Printf("取消订单：%v 时间：%v \n", structName, time.Now().String())
}

func (w *Waiter) NotifyCommand() {
	for i, cmd := range w.orders {
		println("No. ", i)
		cmd.executeCommand()
	}
}
