package design_pattern

import (
	"fmt"
	"testing"
)

func TestSimpleFactory(t *testing.T)  {
	var opt Operation
	opt, err := CreateOperation("*")
	opt.SetNum(1,6)
	if err == nil {
		fmt.Println(opt.GetResult())
	}
}

func TestStrategy(t *testing.T)  {
	var cashContext CashContext
	cashContext.GetCashStrategy("arrival 300 minus 100")
	totalPrice := cashContext.super.AcceptCash(800)
	fmt.Println(totalPrice)
	cashContext.GetCashStrategy("twenty per cent discount")
	totalPrice = cashContext.super.AcceptCash(800)
	fmt.Println(totalPrice)
}


func TestDecorator(t *testing.T){
	var person Person
	basePerson := &BasePerson{
		Name:"小菜",
	}
	person = NewTie(basePerson)
	person = NewBigTrouser(person)

	person.Show()

}

func TestDecorator2(t *testing.T)  {
	var person Person
	person = &BasePerson{
		Name:"小菜",
	}
	tie := &Tie2{}

	big := &BigTrouser2{}

	person = tie.Decorate(person)

	person = big.Decorate(person)
	person.Show()

}

func TestDecoratorLog(t *testing.T)  {
	Log("msg", "first")
}

func TestProxy(t *testing.T)  {
	persuit := &Persuit{
		SchoolGirl:"jiaojiao",
	}
	proxy := &Proxy{
		*persuit,
	}

	proxy.GiveDolls()
	proxy.GiveFlowers()
	proxy.GiveChocolate()
}

func TestFactoryMethoh(t *testing.T)  {
	iFactory := &UndergraduateFactory{}
	ugleifeng := iFactory.CreateLeiFeng()
	ugleifeng.Sweep()
	ugleifeng.BuyRice()
	ugleifeng.Wash()
}