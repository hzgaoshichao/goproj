package main

type DepartmentInterface interface {
	insert(department Department)
	getDepartment(id int) *Department
}

type Department struct {
	id   int
	name string
}

func (u *Department) SetId(id int) {
	u.id = id
}

func (u *Department) GetId(id int) int {
	return u.id
}

func (u *Department) SetName(name string) {
	u.name = name
}

func (u *Department) GetName(name string) string {
	return u.name
}
