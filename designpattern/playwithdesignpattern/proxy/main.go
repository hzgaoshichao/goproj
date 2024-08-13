package main

func main() {
	sg := SchoolGirl{}
	sg.setName("小美")

	p := newProxy(&sg)
	p.giveDolls()
	p.giveFlowers()
	p.giveChocolate()

}
