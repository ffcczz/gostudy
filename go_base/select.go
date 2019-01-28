package go_base

import (
	"fmt"
)

type Task struct {
	task chan chan int
	quit chan struct{}
}

func (t *Task)TaskTest()  {
	taskC := make(chan int)
	fmt.Println(t)

	select {
	case t.task <- taskC:
		fmt.Println("taskC inside chan")
	case <-t.quit:
		fmt.Println(" first quit")
	}

	select {
	case taskC <- 1:
		fmt.Println("1 inside taskC chan")
	case <- t.quit:
		fmt.Println("quit")
	}


}