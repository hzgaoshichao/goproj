package main

type SubjectInterface interface {
	attach(observer ObserverInterface)
	detach(observer ObserverInterface)
	notifyEmployee()
	getSubject() *Subject // 由于接口中不能定义字段, 这个方法用于返回公共的字段结构体
}

type Subject struct {
	name      string
	observers []ObserverInterface
	action    string
}

func (s *Subject) attach(observer ObserverInterface) {
	s.observers = append(s.observers, observer)
}

func (s *Subject) detach(observer ObserverInterface) {
	// TODO
}

func (s *Subject) notifyEmployee() {
	for _, observer := range s.observers {
		observer.update()
	}
}

func (s *Subject) getAction() string {
	return s.action
}

func (s *Subject) setAction(action string) {
	s.action = action
}
