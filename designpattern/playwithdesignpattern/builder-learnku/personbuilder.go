package main

type PersonBuilder struct {
	person *Person
}

func (p *PersonBuilder) WithName(name string) *PersonBuilder {
	p.person.name = name
	return p
}

func (p *PersonBuilder) WithAge(age int) *PersonBuilder {
	p.person.age = age
	return p
}

func (p *PersonBuilder) WithGender(gender string) *PersonBuilder {

	p.person.gender = gender
	return p
}

func (p *PersonBuilder) WithAddress(address string) *PersonBuilder {
	p.person.address = address
	return p
}

func (p *PersonBuilder) Build() *Person {
	return p.person
}

func NewPersonBuilder() *PersonBuilder {
	return &PersonBuilder{
		person: &Person{},
	}
}
