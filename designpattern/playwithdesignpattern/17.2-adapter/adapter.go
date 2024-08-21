package main

type adapter struct {
	a adaptee
}

func (a *adapter) request() {
	a.a.specificRequest()
}

func NewAdapter() adapter {
	return adapter{
		a: adaptee{},
	}
}
