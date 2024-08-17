package main

type director struct {
}

func (d *director) construct(builder builderInterface) {
	builder.buildPartA()
	builder.buildPartB()
}
