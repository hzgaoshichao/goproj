package main

import (
	"fmt"
)

type hrDepartment struct {
	name string
}

func NewHRDepartment(name string) *hrDepartment {
	return &hrDepartment{
		name: name,
	}
}

func (h *hrDepartment) add(com companyInterface) {

}

func (h *hrDepartment) remove(com companyInterface) {

}

func (h *hrDepartment) display(depth int) {
	for i := 0; i < depth; i++ {
		fmt.Printf("-")
	}
	fmt.Printf("%v \n", h.name)

}

func (h *hrDepartment) lineOfDuty() {

	fmt.Printf("%v 员工招聘培训管理 \n", h.name)
}
