package test

import (
	"fmt"
	"sort"
	"testing"
)

func Test_sort_string(t *testing.T)  {
	sort_string := sort.StringSlice{
		"test",
		"check",
		"pro",
		"best",
	}
	sort_string.Sort()
	fmt.Println(sort_string)

}
