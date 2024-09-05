package demo

import "fmt"

type Colleague interface {
	Send(message string)
	Notify(message string)
}

type ColleagueCommon struct {
	mediator Mediator
}

func (cc *ColleagueCommon) Send(message string) {
	fmt.Printf("Not Implemented")
}

func (cc *ColleagueCommon) Notify(message string) {
	fmt.Printf("Not Implemented")
}

func NewColleagueCommon(mediator Mediator) *ColleagueCommon {
	return &ColleagueCommon{
		mediator: mediator,
	}
}

type ConcreteColleague1 struct {
	ColleagueCommon
}

func NewConcreteColleague1(mediator Mediator) *ConcreteColleague1 {
	return &ConcreteColleague1{
		ColleagueCommon: *NewColleagueCommon(mediator),
	}
}

func (cc *ConcreteColleague1) Send(message string) {
	cc.ColleagueCommon.mediator.send(message, cc)
}

func (cc *ConcreteColleague1) Notify(message string) {
	fmt.Printf("同事1得到信息: %v \n", message)
}

type ConcreteColleague2 struct {
	ColleagueCommon
}

func NewConcreteColleague2(mediator Mediator) *ConcreteColleague2 {
	return &ConcreteColleague2{
		ColleagueCommon: *NewColleagueCommon(mediator),
	}
}

func (cc *ConcreteColleague2) Send(message string) {
	cc.ColleagueCommon.mediator.send(message, cc)
}

func (cc *ConcreteColleague2) Notify(message string) {
	fmt.Printf("同事2得到信息: %v \n", message)
}
