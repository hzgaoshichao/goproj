package main

import "fmt"

type Pursuit struct {
	schoolgirl *SchoolGirl
}

func newPursuit(sg *SchoolGirl) *Pursuit {
	return &Pursuit{
		schoolgirl: sg,
	}
}

func (p *Pursuit) giveDolls() {
	fmt.Printf("%v, 你好, 送你洋娃娃", p.schoolgirl.name)
}

func (p *Pursuit) giveFlowers() {
	fmt.Printf("%v, 你好, 送你鲜花", p.schoolgirl.name)
}

func (p *Pursuit) giveChocolate() {
	fmt.Printf("%v, 你好, 送你巧克力", p.schoolgirl.name)
}
