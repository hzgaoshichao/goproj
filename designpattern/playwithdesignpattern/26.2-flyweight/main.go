package main

import d "flyweight/demo"

func main() {
	extrinsicstate := int(22)

	f := d.NewFlyweightFactory()
	fx := f.GetFlyweight("X")
	extrinsicstate--
	fx.Operation(extrinsicstate)
	fy := f.GetFlyweight("Y")
	extrinsicstate--
	fy.Operation(extrinsicstate)
	fz := f.GetFlyweight("Z")
	extrinsicstate--
	fz.Operation(extrinsicstate)

	uf := d.UnsharedConcreteFlyweight{}
	extrinsicstate--
	uf.Operation(extrinsicstate)
}
