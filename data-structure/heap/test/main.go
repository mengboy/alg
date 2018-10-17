package main

import (
	"alg/data-structure/heap"
	"fmt"
)

func main()  {
	//h := heap.NewMin()
	//h.Insert(3)
	//h.Insert(5)
	//h.Insert(2)
	//h.Insert(6)
	//h.Insert(4)
	//h.Insert(1)
	//fmt.Println(h.Data)
	//
	//x := make([]int, 0)
	//
	//for h.Len > 0{
	//	x = append(x, *h.Pop())
	//	fmt.Println(h.Data)
	//}
	//
	//fmt.Println(x)

	h2 := heap.NewMin()
	h2.Data = []int{
		3, 5, 6, 2, 1, 8, 9,
	}

	h2.Len = 7
	//h2.ShiftUp()
	h2.Heapfiy()
	fmt.Println(h2.Data)
	//
	//HeapSort([]int{
	//		3, 5, 6, 2, 1, 8, 9,
	//	})

}


func HeapSort(array []int) {
	m := len(array)
	s := m/2
	for i := s; i > -1; i-- {
		heapfiy(array, i, m-1)
	}
	fmt.Println(array)
	for i := m-1; i > 0; i-- {
		array[i], array[0] = array[0], array[i]
		heapfiy(array, 0, i-1)
	}
	fmt.Println(array)
}

func heapfiy(array []int, i, end int){
	l := 2*i+1
	if l > end {
		return
	}
	n := l
	r := 2*i+2
	if r <= end && array[r]>array[l]{
		n = r
	}
	if array[i] > array[n]{
		return
	}
	array[n], array[i] = array[i], array[n]
	heapfiy(array, n, end)
}
