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
// 冒泡排序2
func (sorter *Sorter)SortBubble2(array []int)  {
	for i := 0; i < len(array); i++ {
		for j := i+1; j < len(array); j++{
			if array[j-1] > array[j] {
				array[j-1],array[j] = array[j],array[j-1]
			}
		}

	}
}
// 冒泡排序3
func (sorter *Sorter)SortBubble3(array []int)  {
	for i := 0; i < len(array); i++ {
		for j := len(array)-1; j > i; j-- {
			if array[j-1] > array[j] {
			//if array[j] > array[j-1] {
				array[j-1],array[j] = array[j],array[j-1]
			}
		}
	}
}
// 冒泡排序4--改进冒泡排序
func (sorter *Sorter) SortBubble4(array []int)  {
	flag := true
	for i := 0; i < len(array) && flag; i++ {
		flag = false
		for j := len(array)-1; j > i; j--{
			if array[j-1] > array[j]{
				array[j-1],array[j] = array[j],array[j-1]
				flag = true   // 数据有位置交换,则循环继续,数据没有交换,则循环终止
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

// 选择排序again  相比于冒泡排序，选择排序的交换次数要少
func (sorter *Sorter) SortSelect2(array []int)  {
	for i := 0; i < len(array); i++ {
		min := i
		for j := i+1 ; j < len(array); j++{
			if array[min] > array[j]{
				min = j
			}
		}
		if min != i {
			array[min],array[i] = array[i],array[min]
		}
	}
}

// 直接插入排序
func (sorter *Sorter)SortInsert(array  []int)  {
	arraylength  := len(array)
	for i, j := 1,0; i < arraylength; i++ {
		temp := array[i]
		for j = i; j > 0 && array[j-1] > temp; j-- {
			array[j] = array[j-1]
			fmt.Printf("i=%d,j=%d,array=%v\n", i, j, array)
		}
		array[j] = temp
	}
}

// 直接插入排序2
func (sorter *Sorter) SortInsert2(array []int)  {
	for i,j := 1,0; i < len(array); i++ {
		if array[i-1] > array[i]{
			//此内层循环在于将大于array[i]的数往后移
			temp := array[i]
			for j = i-1; j >= 0 && array[j] > temp; j-- {
				array[j+1] = array[j]
			}
			array[j+1] = temp
		}
	}
}


func main()  {
	//冒泡排序
	array := []int{6, 4, 7, 22, 19, 27, 32, 91,54, }
	learnsort := Sorter{name: "冒泡排序"}
	learnsort.SortBubble(array)
	fmt.Println(learnsort.name, array)

	//冒泡排序2
	array = []int{6, 4, 7, 22, 19, 27, 32, 91,54, }
	learnsort2 := Sorter{name: "冒泡排序2"}
	learnsort2.SortBubble2(array)
	fmt.Println(learnsort2.name, array)

	//冒泡排序3
	array = []int{6, 4, 7, 22, 19, 27, 32, 91,54, }
	learnsort3 := Sorter{name: "冒泡排序3"}
	learnsort3.SortBubble3(array)
	fmt.Println(learnsort3.name, array)

	//冒泡排序4
	array = []int{6, 4, 7, 22, 19, 27, 32, 91,54, }
	learnsort4 := Sorter{name: "冒泡排序4"}
	learnsort4.SortBubble4(array)
	fmt.Println(learnsort4.name, array)

	//快速排序
	array = []int{6, 4, 7, 22, 19, 27, 32, 91,54, }
	repidsort := Sorter{name: "快速排序"}
	repidsort.SortRapid(array)
	fmt.Println(repidsort.name, array)

	//选择排序
	array = []int{6, 4, 7, 22, 19, 27, 32, 91,54, }
	selectsort := Sorter{name: "选择排序"}
	selectsort.SortSelect(array)
	fmt.Println(selectsort.name, array)

	//选择排序2
	array = []int{6, 4, 7, 22, 19, 27, 32, 91,54, }
	selectsort2 := Sorter{name: "选择排序2"}
	selectsort2.SortSelect2(array)
	fmt.Println(selectsort2.name, array)

	//插入排序
	array = []int{6, 4, 7, 22, 19, 27, 32, 91,54, }
	insertsort := Sorter{name : "插入排序"}
	insertsort.SortInsert(array)
	fmt.Println(insertsort.name, array)

	//插入排序2
	array = []int{6, 4, 7, 22, 19, 27, 32, 91,54, }
	insertsort2 := Sorter{name : "插入排序2"}
	insertsort2.SortInsert2(array)
	fmt.Println(insertsort2.name, array)
}