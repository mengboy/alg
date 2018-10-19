package main

import (
	"alg/data-structure/stack"
	"fmt"
)

func main()  {
	arr := []int{
		20, 1, 2, 3, 7, 4, 9, 10, 19, 17, 15,
	}
	//QuickSort(arr, 0, len(arr) - 1)
	//QuickSort2(arr)
	//quickSort(arr)
	qsort(arr)
	fmt.Println(arr)
}


// 递归
func quickSort(arr []int)  {
	if len(arr) <= 1 {
		return
	}
	mid := arr[0]
	head, tail := 0, len(arr) - 1
	for i := 1; i <= tail; {
		// 将mid左侧所有大于mid的数放到右侧末尾
		if arr[i] > mid{
			arr[i], arr[tail] = arr[tail], arr[i]
			tail--
		}else {
			// 小于mid的数放到mid左侧开始处
			arr[i], arr[head] = arr[head], arr[i]
			head++
			i++
		}
	}
	quickSort(arr[:head])
	quickSort(arr[head+1:])
}

func qsort(data []int) {
	if len(data) <= 1 {
		return
	}
	mid := data[0]
	head, tail := 0, len(data)-1
	for i := 1; i <= tail; {
		if data[i] > mid {
			data[i], data[tail] = data[tail], data[i]
			tail--
		} else {
			data[i], data[head] = data[head], data[i]
			head++
			i++
		}
	}
	qsort(data[:head])
	qsort(data[head+1:])
}


// 递归
func QuickSort(arr []int, left, right int)  {
	if left < right{
		i := partition(arr, left, right)
		QuickSort(arr, left, i)
		QuickSort(arr, i + 1, right)
	}
}

func partition(arr []int, left, right int) int{
	mid := arr[left]
	i, j := left, right
	for i < j {
		// 从右边找到比mid小的数
		for ; i < j && arr[j] >= mid; j--  {
		}
		arr[i] = arr[j]
		// 从左边找到比mid大的数
		for ; i < j && arr[i] < mid; i++{
		}
		arr[j] = arr[i]
	}
	// 恢复
	arr[i] = mid
	return i
}


// 非递归

func QuickSort2(arr []int)  {
	// 借助栈将需要partition的数据先压入栈中
	s := stack.New()
	l := 0
	r := len(arr) - 1
	s.Push(l)
	s.Push(r)
	for !s.Empty(){
		right := s.Pop().(int)
		left := s.Pop().(int)
		if left >= right{
			continue
		}
		i := partition(arr, left, right)
		s.Push(left)
		s.Push(i)
		s.Push(i + 1)
		s.Push(right)
	}

}