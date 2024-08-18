package main

func main() {

	boss := NewBoss("胡汉三")
	employee1 := NewStockObserver("魏关姹", boss)
	employee2 := NewStockObserver("易管查", boss)
	employeelqm := NewNBAObserver("兰秋冥", boss)
	emphardworker := NewNBAObserver("打工人", boss)

	boss.attach(employee1)
	boss.attach(employee2)
	boss.attach(emphardworker)
	boss.attach(employeelqm)
	boss.detach(emphardworker)
	boss.setAction("我胡汉三回来了")
	boss.notifyEmployee()

	secretary01 := NewSecretary("秘书A")
	employee3 := NewStockObserver("魏关姹", secretary01)
	secretary01.attach(employee3)
	secretary01.setAction("老板回来了")
	secretary01.notifyEmployee()

}
