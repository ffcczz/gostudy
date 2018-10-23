package test

import (
	"fmt"
	"konggostudy/struct"
	"testing"
)

func Test(t *testing.T)  {
	wheel := &_struct.Wheel{}
	wheel.X = 10
	fmt.Printf("%#v", wheel)
}
