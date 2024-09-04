package main

import bbq "cmdpattern/barbecue"

func main() {
	boy := bbq.Barbecuer{}
	bakeMuttonCommand1 := bbq.NewBakeMuttonCommand(&boy)
	bakeChickenWindCommand1 := bbq.NewBakeChickenWingCommand(&boy)
	girl := bbq.Waiter{}
	girl.SetOrder(bakeMuttonCommand1)
	girl.NotifyCommand()
	girl.SetOrder(bakeChickenWindCommand1)
	girl.NotifyCommand()
}
