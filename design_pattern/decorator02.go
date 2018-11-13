package design_pattern

import "fmt"

type Finery interface {
	Show()
	Decorate(Person) Person
}

type Tshirts2 struct {
	person Person
}

func (t *Tshirts2)Decorate(person Person) Person {
	return &Tshirts2{person}
}

func (t *Tshirts2) Show()  {
	fmt.Print("大T恤")
	t.person.Show()
}

type BigTrouser2 struct {
	person Person
}

func (bt *BigTrouser2)Decorate(person Person) Person {
	return &BigTrouser2{person}
}

func (bt *BigTrouser2) Show()  {
	fmt.Print("大裤衩")
	bt.person.Show()
}

type Tie2 struct {
	Name string
	person Person
}

func (t *Tie2)Decorate(person Person) Person {
	return &Tie2{
		Name:"张三",
	person:person}
}

func (t *Tie2) Show()  {
	fmt.Print("领带:" + t.Name)
	t.person.Show()
}



