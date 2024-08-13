package main

type SchoolGirl struct {
	name string
}

func (sg *SchoolGirl) setName(name string) {
	sg.name = name
}
