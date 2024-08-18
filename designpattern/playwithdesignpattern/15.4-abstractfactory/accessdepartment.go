package main

import "fmt"

type AccessDepartment struct {
}

func (s *AccessDepartment) insert(department Department) {
	fmt.Printf("在 Access 中给 Department 增加了一条记录 \n")
}
func (s *AccessDepartment) getDepartment(id int) *Department {
	fmt.Printf("在 Access 中根据用户 id 得到 Department 表的一条记录 \n")
	return nil
}
