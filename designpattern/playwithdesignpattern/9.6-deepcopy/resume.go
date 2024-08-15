package main

import "fmt"

type resume struct {
	name string
	sex  string
	age  string
	work workExperience
}

func NewResume(name string) resume {
	return resume{
		name: name,
		work: workExperience{},
	}

}

func (r *resume) setPersonalInfo(sex string, age string) {
	r.sex = sex
	r.age = age
}

func (r *resume) setWorkExperience(timeArea string, company string) {
	r.work.setTimeArea(timeArea)
	r.work.setCompany(company)

}

func (r *resume) display() {
	fmt.Printf("name: %v, sex: %v, age: %v \n", r.name, r.sex, r.age)
	fmt.Printf("工作经历 %v %v \n", r.work.timeArea, r.work.company)
}

func (r *resume) clone() cloneableInterface {
	newResume := *r
	// newResume.work = r.work
	return &newResume
}
