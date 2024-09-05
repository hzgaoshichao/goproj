package demo

import "reflect"

type Mediator interface {
	send(message string, colleague Colleague)
}

type ConcreteMediator struct {
	colleague1 *ConcreteColleague1
	colleague2 *ConcreteColleague2
}

func (cm *ConcreteMediator) SetColleague1(value *ConcreteColleague1) {
	cm.colleague1 = value
}

func (cm *ConcreteMediator) SetColleague2(value *ConcreteColleague2) {
	cm.colleague2 = value
}

func (cm *ConcreteMediator) send(message string, colleague Colleague) {
	if reflect.DeepEqual(cm.colleague1, colleague) {
		cm.colleague2.Notify(message)
	} else {
		cm.colleague1.Notify(message)
	}
}
