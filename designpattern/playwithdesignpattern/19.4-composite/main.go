package main

import "fmt"

func main() {
	root := NewConcreteCompany("北京总公司")
	root.add(NewFinanceDepartment("总公司财务部"))
	root.add(NewHRDepartment("总公司人力资源"))

	comp := NewConcreteCompany("上海华东分公司")
	comp.add(NewHRDepartment("华东分公司人力资源部"))
	comp.add(NewFinanceDepartment("华东分公司财务部"))
	root.add(comp)

	comp2 := NewConcreteCompany("南京办事处")
	comp2.add(NewHRDepartment("南京办事处人力资源部"))
	comp2.add(NewFinanceDepartment("南京办事处财务部"))
	comp.add(comp2)

	comp3 := NewConcreteCompany("杭州办事处")
	comp3.add(NewHRDepartment("杭州办事处人力资源部"))
	comp3.add(NewFinanceDepartment("杭州办事处财务部"))
	comp.add(comp3)

	fmt.Printf("结构图: \n")
	root.display(1)
	fmt.Printf("职责: \n")
	root.lineOfDuty()
}
