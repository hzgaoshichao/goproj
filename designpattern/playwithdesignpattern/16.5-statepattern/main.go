package main

func main() {
	c := NewContext(&ConcreteStateA{})
	c.request()
	c.request()
	c.request()
	c.request()

}
