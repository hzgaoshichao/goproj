package main

type Proxy struct {
	pursuit *Pursuit
}

func newProxy(sg *SchoolGirl) *Proxy {
	return &Proxy{
		pursuit: newPursuit(sg),
	}
}
func (p *Proxy) giveDolls() {
	p.pursuit.giveDolls()
}

func (p *Proxy) giveFlowers() {
	p.pursuit.giveFlowers()
}

func (p *Proxy) giveChocolate() {
	p.pursuit.giveChocolate()
}
