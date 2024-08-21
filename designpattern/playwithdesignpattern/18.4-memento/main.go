package main

func main() {
	o := origator{}
	o.setState("On")
	o.show()

	c := caretaker{}
	c.setMemento(o.createMemento())

	o.setState("OFF")
	o.show()

	o.recoverMemento(c.getMemento())
	o.show()
}
