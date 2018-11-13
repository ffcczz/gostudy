package design_pattern

import "fmt"

type GiveGift interface {
	GiveDolls()
	GiveFlowers()
	GiveChocolate()
}

type Persuit struct {
	SchoolGirl string
}

func (p *Persuit) GiveDolls() {
	fmt.Println("give girl dolls:"+p.SchoolGirl)
}

func (p *Persuit) GiveFlowers() {
	fmt.Println("give girl flowers:" + p.SchoolGirl)
}

func (p *Persuit) GiveChocolate() {
	fmt.Println("give girl chocolate:"+p.SchoolGirl)
}

type Proxy struct {
	persuit Persuit
}

func (p *Proxy) GiveDolls() {
	fmt.Println("I'm proxy")
	p.persuit.GiveDolls()
}

func (p *Proxy) GiveFlowers() {
	fmt.Println("I'm proxy")
	p.persuit.GiveFlowers()
}

func (p *Proxy) GiveChocolate() {
	fmt.Println("I'm proxy")
	p.persuit.GiveChocolate()
}

