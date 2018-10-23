package main

import (
	"errors"
	"fmt"
	"github.com/jinzhu/gorm"
	"time"
)

func main()  {
	TestTimeString()
	TestTimePointer()
}

func TestTimeString()  {
	ti := time.Now()
	fmt.Println(ti)   //2018-10-23 10:35:36.450708984 +0800 CST m=+0.001759227
}


func TestTimePointer()  {
	type timetest  struct {
		ti time.Time
	}
	titest := timetest{ti: time.Time{}}
	fmt.Println(titest.ti)
}

func TestError()  {
	err := errors.New("record not found")
	if err.Error() == gorm.ErrRecordNotFound.Error() {
		fmt.Println("true")
	} else {
		fmt.Println("false")
	}
}