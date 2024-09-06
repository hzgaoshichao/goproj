package main

import un "mediator/unitednation"

func main() {
	unsc := un.UnitedNationsSecurityCouncil{}
	c1 := un.NewUSA(&unsc)
	c2 := un.NewIrap(&unsc)
	unsc.SetUSA(c1)
	unsc.SetIrap(c2)

	c1.Declare("不准研制核武器, 否则要发动战争!\n")
	c2.Declare("我们没有核武器,也不怕侵略! \n")

}
