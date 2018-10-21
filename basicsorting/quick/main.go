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
	//qsort(arr)
	t := quick(arr, 0, len(arr) - 1, 5)
	if t != -1{
		fmt.Println(arr, arr[t-1])
	}

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
		for ; i < j && arr[j] <= mid; j--  {
		}
		arr[i] = arr[j]
		// 从左边找到比mid大的数
		for ; i < j && arr[i] > mid; i++{
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

// partition
func partition2(arr []int, l, r int, k int)  int {
	i := l
	j := r
	mid := arr[l]
	for i < j{
		for i < j && arr[i] < mid{
			i++
		}
		// 将大数移到后面
		arr[j] = arr[i]
		for i < j && arr[j] > mid{
			j--
		}
		// 将小数移到前面
		arr[i] = arr[j]
	}
	// 将分割标准移到中间
	arr[i] = mid
	return i
}

// 递归
func quick(arr []int, l, r int, k int) int {
	if l > r{
		return -1
	}
	i := partition(arr, l, r)
	if i == k{
		return k
	}
	if i < k{
		return quick(arr, i + 1, r, k)
	}else {
		return quick(arr, l, i, k)
	}

	return -1
}

// 非递归
func quickNo(arr []int, l, r int, k int)  {
	s := stack.New()
	s.Push(l)
	s.Push(r)
	for !s.Empty(){
		end := s.Pop().(int)
		start := s.Pop().(int)
		if start >= end{
			continue
		}
		i := partition2(arr, start, end, k)
		s.Push(start)
		s.Push(i)
		s.Push(i + 1)
		s.Push(end)
	}
}