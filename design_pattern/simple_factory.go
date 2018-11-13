package design_pattern

import "errors"

type Operation interface {
	SetNum(Num1, Num2 int)
	GetResult() int
}

type BaseOperation struct {
	Num1 int
	Num2 int
}

func (base *BaseOperation) SetNum(Num1, Num2 int)  {
	base.Num1 = Num1
	base.Num2 = Num2
}

type OperationAdd struct{
	BaseOperation
}

func (add *OperationAdd)GetResult() int  {
	return add.Num1 + add.Num2
}

type OperationSub struct {
	BaseOperation
}

func (sub *OperationSub) GetResult() int {
	return sub.Num1 - sub.Num2
}

type OperaiontMul struct {
	BaseOperation
}

func (mul *OperaiontMul) GetResult() int {
	return mul.Num1 * mul.Num2
}

type OperationDiv struct {
	BaseOperation
}

func (div *OperationDiv) GetResult() int {
	return div.Num1 / div.Num2
}

func CreateOperation(operation string)(ope Operation, err error)  {
	switch operation {
	case "+":
		ope = &OperationAdd{}
	case "-":
		ope = &OperationSub{}
	case "*":
		ope = &OperaiontMul{}
	case "/":
		ope = &OperationDiv{}
	default:
		err = errors.New("运算符错误")
	}
	return
}


