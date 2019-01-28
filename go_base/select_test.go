package go_base

import (
	"testing"
	"time"
	"fmt"
)

func TestTask_TaskTest(t *testing.T) {
	ta := Task{
		task:make(chan chan int),
		quit:make(chan struct{}),
	}

	go func() {
		ta.quit <- struct{}{}}()

	go ta.TaskTest()
	go ta.TaskTest()

	tas := <- ta.task
	fmt.Println("1 outside chan", <- tas)



	time.Sleep(time.Second * 3)


}
