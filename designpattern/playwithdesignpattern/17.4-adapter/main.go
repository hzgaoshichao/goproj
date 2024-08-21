package main

func main() {
	forward := NewForward("巴蒂尔")
	forward.attack()
	forward.defense()

	guard := NewGuards("麦克格雷迪")
	guard.attack()
	guard.defense()

	center := NewTranslator("姚明")
	center.attach()
	center.defense()

}
