package main

import (
	bbq "cmdpattern/barbecue"
	"fmt"
)

func main() {
	boy := bbq.Barbecuer{}
	bakeMuttonCommand1 := bbq.NewBakeMuttonCommand(&boy)
	bakeChickenWindCommand1 := bbq.NewBakeChickenWingCommand(&boy)

	girl := bbq.NewWaiter()
	fmt.Printf("开门营业， 顾客点菜\n")
	girl.SetOrder(bakeMuttonCommand1)
	girl.SetOrder(bakeMuttonCommand1)
	girl.SetOrder(bakeMuttonCommand1)
	girl.SetOrder(bakeMuttonCommand1)
	girl.SetOrder(bakeMuttonCommand1)
	girl.SetOrder(bakeMuttonCommand1)
	girl.SetOrder(bakeMuttonCommand1)
	girl.SetOrder(bakeMuttonCommand1)
	girl.CancelOrder(bakeMuttonCommand1)
	girl.SetOrder(bakeChickenWindCommand1)

	fmt.Printf("点菜完毕， 通知厨房烧菜\n")
	girl.NotifyCommand()
}
