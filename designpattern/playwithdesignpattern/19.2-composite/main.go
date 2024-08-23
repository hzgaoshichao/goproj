package main

func main() {
	root := NewComposite("root")
	root.add(NewLeaf("LeafA"))
	root.add(NewLeaf("LeafB"))

	comp := NewComposite("Composite X")
	comp.add(NewLeaf("LeafXA"))
	comp.add(NewLeaf("LeafXB"))

	root.add(comp)

	comp2 := NewComposite("Composite XY")
	comp2.add(NewLeaf("LeafXYA"))
	comp2.add(NewLeaf("LeafXYB"))

	comp.add(comp2)

	leaf := NewLeaf("Leaf C")
	root.add(leaf)

	leaf2 := NewLeaf("Leaf D")
	root.add(leaf2)
	root.remove(leaf2)

	root.display(10)
}
