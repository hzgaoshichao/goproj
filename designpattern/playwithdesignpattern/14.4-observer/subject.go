package main

import "reflect"

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
	for i, ob := range s.observers {
		//if observer == ob { // 这里涉及到结构体的比较, 需要特别注意 // 参考链接: https://segmentfault.com/a/1190000040099215
		if reflect.DeepEqual(observer, ob) {
			s.observers = append(s.observers[:i], s.observers[i+1:]...)
		}
	}
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
