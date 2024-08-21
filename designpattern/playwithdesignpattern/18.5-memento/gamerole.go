package main

import "fmt"

type gamerole struct {
	vitality int
	attack   int
	defense  int
}

func (g *gamerole) getVitality() int {
	return g.vitality
}

func (g *gamerole) setVitality(value int) {
	g.vitality = value
}

func (g *gamerole) getAttack() int {
	return g.attack
}

func (g *gamerole) setAttack(value int) {
	g.attack = value
}

func (g *gamerole) getDefense() int {
	return g.defense
}

func (g *gamerole) setDefense(value int) {
	g.defense = value
}

func (g *gamerole) displayState() {
	fmt.Printf("角色当前状态: \n")
	fmt.Printf("体力: %v \n", g.vitality)
	fmt.Printf("攻击力: %v \n", g.attack)
	fmt.Printf("防御力: %v \n", g.defense)
}

func (g *gamerole) getInitState() {
	g.vitality = 100
	g.attack = 100
	g.defense = 100
}

func (g *gamerole) fight() {
	g.vitality = 0
	g.attack = 0
	g.defense = 0
}

func (g *gamerole) saveState() *roleStateMemento {

	return NewRoleStateMemento(g.vitality, g.attack, g.defense)

}

func (g *gamerole) recoverState(rsm *roleStateMemento) {
	g.vitality = rsm.vitality
	g.attack = rsm.attack
	g.defense = rsm.defense
}
