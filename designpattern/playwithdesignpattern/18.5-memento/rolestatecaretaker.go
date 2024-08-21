package main

type roleStateMementoCaretaker struct {
	rsm *roleStateMemento
}

func (r *roleStateMementoCaretaker) getRoleStateMemento() *roleStateMemento {
	return r.rsm
}
func (r *roleStateMementoCaretaker) setRoleStateMemento(value *roleStateMemento) {
	r.rsm = value
}
