package main

type UserInterface interface {
	insert(user User)
	getUser(id int) *User
}

type User struct {
	id   int
	name string
}

func (u *User) SetId(id int) {
	u.id = id
}

func (u *User) GetId(id int) int {
	return u.id
}

func (u *User) SetName(name string) {
	u.name = name
}

func (u *User) GetName(name string) string {
	return u.name
}
