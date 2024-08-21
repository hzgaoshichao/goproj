package main

type memento struct {
	state string
}

func NewMemento(state string) *memento {
	return &memento{
		state: state,
	}
}

func (m *memento) getState() string {
	return m.state
}

func (m *memento) setState(value string) {
	m.state = value
}
