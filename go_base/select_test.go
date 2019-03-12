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
	//ta.TaskTest()

	/*go func() {
		time.Sleep(10 * time.Second)
		ta.quit <- struct{}{}}()*/

	go ta.TaskTest()
	//go ta.TaskTest()


	go func() {
		time.Sleep(5 * time.Second)
		tas := <- ta.task
		fmt.Println("1 outside chan", <- tas)
	}()




	time.Sleep(time.Second * 5)
	select{}


}
