package design_pattern

import "fmt"


type Person interface {
	Show()
}

/*type Finary interface {
	Show()
	Decorate(person BasePerson)
}*/

type BasePerson struct {
	Name string
}

func (base *BasePerson) Show() {
	fmt.Println("Name:", base.Name)
}

type Tshirts struct {
	person Person
}

func NewTshirts(person Person) *Tshirts {
	return &Tshirts{person}
}

func (t *Tshirts) Show()  {
	fmt.Print("大T恤")
	t.person.Show()
}

type BigTrouser struct {
	person Person
}

func NewBigTrouser(person Person) *BigTrouser {
	return &BigTrouser{person}
}

func (bt *BigTrouser) Show()  {
	fmt.Print("大裤衩")
	bt.person.Show()
}

type Tie struct {
	person Person
}

func NewTie(person Person) *Tie {
	return &Tie{person}
}

func (t *Tie) Show()  {
	fmt.Print("领带")
	t.person.Show()
}



