package test

import (
	"fmt"
	"testing"
	"gostudy/struct"
)

func Test(t *testing.T)  {
	wheel := &_struct.Wheel{}
	wheel.X = 10
	fmt.Printf("%#v", wheel)
}
