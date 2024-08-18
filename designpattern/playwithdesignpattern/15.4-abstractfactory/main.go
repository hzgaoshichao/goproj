package main

func main() {
	user := User{}
	factory := AccessFactory{} // SqlserverFactory{}
	iu := factory.CreateUser()
	iu.insert(user)
	iu.getUser(1)

	dp := Department{}
	idep := factory.CreateDepartment()
	idep.insert(dp)
	idep.getDepartment(1)
}
