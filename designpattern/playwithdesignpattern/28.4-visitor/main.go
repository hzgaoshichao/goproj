package main

import pa "visitor/personaction"

func main() {
	o := pa.NewObjectStructure()
	o.Attach(&pa.Man{})
	o.Attach(&pa.Woman{})

	v1 := pa.Success{}
	o.Display(&v1)

	v2 := pa.Failing{}
	o.Display(&v2)

	v3 := pa.Amativeness{}
	o.Display(&v3)

	v4 := pa.Marriage{}
	o.Display(&v4)
}
