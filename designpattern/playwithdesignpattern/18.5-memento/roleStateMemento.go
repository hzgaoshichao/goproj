package main

type roleStateMemento struct {
	vitality int
	attack   int
	defense  int
}

func NewRoleStateMemento(vitality, attach, defense int) *roleStateMemento {
	return &roleStateMemento{
		vitality: vitality,
		attack:   attach,
		defense:  defense,
	}
}

func (r *roleStateMemento) getVitality() int {
	return r.vitality
}

func (r *roleStateMemento) getAttack() int {
	return r.attack
}

func (r *roleStateMemento) getDefense() int {
	return r.defense
}
