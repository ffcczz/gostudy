package test

import (
	"testing"
	"fmt"
)

func TestBreak(t *testing.T)  {
	flag:
		for i := 0;i < 10; i++ {
			if i == 5 {
				break flag
			}
			fmt.Println("Testbreak", i)
		}

		fmt.Println("break")
}