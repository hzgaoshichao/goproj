package main

import "fmt"

type SqlserverDepartment struct {
}

func (s *SqlserverDepartment) insert(department Department) {
	fmt.Printf("在 Sql Server 中给 Department 增加了一条记录 \n")
}
func (s *SqlserverDepartment) getDepartment(id int) *Department {
	fmt.Printf("在 Sql Server 中根据用户 id 得到 Department 表的一条记录 \n")
	return nil
}
