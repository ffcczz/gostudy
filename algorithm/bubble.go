package main

import "fmt"


type Sorter struct {
	name string
}

//冒泡排序
func (sorter *Sorter) SortBubble(array []int)  {
	for i := 0; i < len(array); i ++ {
		for j := 0; j < len(array)-i-1; j++ {
			if array[j] > array[j+1] {
				array[j], array[j+1] = array[j+1], array[j]
			}
		}
	}
}

//快速排序
func (sorter *Sorter) SortRapid(array []int)  {
	if len(array) <= 1 {
		return
	}
	mid := array[0]
	i := 1
	head, tail := 0, len(array)-1
	for head < tail {
		if array[i] > mid {
			array[i], array[tail] = array[tail], array[i]
			tail--
		} else {
			array[i], array[head] = array[head], array[i]
			head++
			i++
	    }
	}

	array[head] = mid
	sorter.SortRapid(array[:head])
	sorter.SortRapid(array[head+1:])

}

//选择排序
func (sorter *Sorter) SortSelect(array []int)  {
	arraylength := len(array)
	for i := 0; i < arraylength; i++ {
		min := i
		for j := i +1; j < arraylength; j++ {
			if array[j] < array[min] {
				min = j
			}
		}
		array[i], array[min] = array[min], array[i]
	}
}
// 插入排序
func (sorter *Sorter)SortInsert(array  []int)  {
	arraylength  := len(array)
	for i, j := 1,0; i < arraylength; i++ {
		temp := array[i]
		for j = i; j > 0 && array[j-1] > temp; j-- {
			array[j] = array[j-1]
			fmt.Printf("i=%d,j=%d,array=%v", i, j, array)
		}
		array[j] = temp
	}
}

func main()  {
	array := []int{6, 4, 7, 22, 19, 27, 32, 54, 91}
	learnsort := Sorter{name: "冒泡排序"}
	learnsort.SortBubble(array)
	fmt.Println(learnsort.name, array)

	array = []int{6, 4, 7, 22, 19, 27, 32, 54, 91}
	repidsort := Sorter{name: "快速排序"}
	repidsort.SortRapid(array)
	fmt.Println(repidsort.name, array)

	array = []int{6, 4, 7, 22, 19, 27, 32, 54, 91}
	selectsort := Sorter{name: "选择排序"}
	selectsort.SortSelect(array)
	fmt.Println(selectsort.name, array)

	array = []int{6, 4, 7, 22, 19, 27, 32, 54, 91}
	insertsort := Sorter{name : "插入排序"}
	insertsort.SortInsert(array)
	fmt.Println(insertsort.name, array)
}