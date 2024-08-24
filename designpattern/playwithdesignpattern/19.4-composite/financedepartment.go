package main

import "fmt"

type financeDepartment struct {
	name string
}

func NewFinanceDepartment(name string) *financeDepartment {
	return &financeDepartment{name: name}
}

func (f *financeDepartment) add(com companyInterface) {

}

func (f *financeDepartment) remove(com companyInterface) {

}

func (f *financeDepartment) display(depth int) {
	for i := 0; i < depth; i++ {
		fmt.Printf("-")
	}
	fmt.Printf("%v \n", f.name)

}

func (f *financeDepartment) lineOfDuty() {

	fmt.Printf("%v 公司财务收支管理 \n", f.name)
}
