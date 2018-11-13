package design_pattern

import (
	"github.com/pkg/errors"
	"math"
)

type CashSuper interface {
	AcceptCash(money float64) float64
}

type CashNormal struct {

}

func (cash *CashNormal) AcceptCash(money float64) float64 {
	return money
}

type CashRebate struct {
	MoneyRebate float64
}


func (cash *CashRebate) AcceptCash(money float64) float64 {
	return money * cash.MoneyRebate
}

type CashReturn struct {
	MoneyCondition float64
	MoneyReturn    float64
}

func (cash *CashReturn) AcceptCash(money float64) float64 {
	if money > cash.MoneyCondition {
		return money - math.Floor(money / cash.MoneyCondition)  * cash.MoneyReturn
	}
	return money
}

type CashContext struct {
	super CashSuper
}

func (cash *CashContext) GetCashStrategy(strategyType string)( err error)  {
	switch strategyType {
	case "normal":
		cash.super = &CashNormal{}
	case "arrival 300 minus 100":
		cash.super = &CashReturn{
			MoneyCondition:300,
			MoneyReturn:100,
		}
	case "twenty per cent discount":
		cash.super = &CashRebate{
			MoneyRebate:0.7,
		}
	default:
		err = errors.New("type error")
	}
	return
}
