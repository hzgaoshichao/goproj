package main

type ObserverInterface interface {
	update()
}

type Observer struct {
	name string
	sub  SubjectInterface
}

func NewObserver(name string, sub SubjectInterface) Observer {
	return Observer{
		name: name,
		sub:  sub,
	}
}
