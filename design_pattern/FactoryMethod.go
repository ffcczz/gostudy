package design_pattern

import "fmt"

type LeiFeng interface {
	Sweep()
	Wash()
	BuyRice()
}

type Undergraduate struct {
}

func (ug *Undergraduate) Sweep() {
	fmt.Println("sweep")
}

func (ug *Undergraduate) Wash() {
	fmt.Println("wash")
}

func (ug *Undergraduate) BuyRice() {
	fmt.Println("buy rice")
}

type  Volunteer struct {

}

func (v *Volunteer) Sweep() {
	fmt.Println("sweep")
}

func (v *Volunteer) Wash() {
	fmt.Println("wash")
}

func (v *Volunteer) BuyRice() {
	fmt.Println("buy rice")
}

type IFactory interface {
	CreateLeiFeng() LeiFeng
}

type UndergraduateFactory struct {
}

func (uf *UndergraduateFactory) CreateLeiFeng() LeiFeng {
	return &Undergraduate{}
}

type VolunteerFactory struct {
}

func (v *VolunteerFactory) CreateLeiFeng() LeiFeng {
	return &Volunteer{}
}

