package main

func main() {
	role := gamerole{}
	role.getInitState()
	role.displayState()

	stateAdmin := roleStateMementoCaretaker{}
	stateAdmin.setRoleStateMemento(role.saveState())

	role.fight()
	role.displayState()

	role.recoverState(stateAdmin.getRoleStateMemento())
	role.displayState()

}
