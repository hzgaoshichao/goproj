package main

type Work struct {
	current      StateInterface
	hour         int
	workFinished bool
}

func NewWork() *Work {
	return &Work{
		current: &ForenoonState{},
	}
}

func (w *Work) SetState(s StateInterface) {
	w.current = s
}

func (w *Work) writeProgram() {
	w.current.writeProgram(w)
}

func (w *Work) getHour() int {
	return w.hour
}

func (w *Work) setHour(h int) {
	w.hour = h
}

func (w *Work) getWorkFinished() bool {
	return w.workFinished
}

func (w *Work) setWorkFinished(value bool) {
	w.workFinished = value
}
