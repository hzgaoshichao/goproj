package main

import "fmt"

type SqlserverUser struct {
}

func (s *SqlserverUser) insert(user User) {
	fmt.Printf("在 Sql Server 中给 User 增加了一条记录 \n")
}
func (s *SqlserverUser) getUser(id int) *User {
	fmt.Printf("在 Sql Server 中根据用户 id 得到 User 表的一条记录 \n")
	return nil
}
