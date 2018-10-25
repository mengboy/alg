package main

import "fmt"

func main()  {
	arr := []int{
		20, 1, 2, 3, 7, 4, 9, 10, 19, 17, 15,
	}
	HeapSort2(arr)
	fmt.Println(arr)
}

func HeapSort(arr []int)  {
	if len(arr) <= 1{
		return
	}
	l := len(arr)
	for i := len(arr) / 2; i >= 0; i--{
		heapfiy(arr, i, l - 1)
	}
	for i := l - 1; i >= 0; i--{
		arr[i], arr[0] = arr[0], arr[i]
		heapfiy(arr, 0, i - 1)
	}
}


// 堆化
// 父 i
// 左 i * 2 + 1
// 右 i * 2 + 2
func heapfiy(arr []int, p int, end int)  {
	l := p * 2 + 1
	if l > end{
		return
	}
	r := p * 2 + 2
	n := l
	if r <= end &&  arr[l] < arr[r]{
		n = r
	}

	if arr[p] > arr[n]{
		return
	}
	arr[p], arr[n] = arr[n], arr[p]
	heapfiy(arr, n, end)
}

func heap(arr []int, p int, end int)  {
	l := p * 2 + 1
	if l > end{
		return
	}
	r := p * 2 + 2
	n := l

	if r < end && arr[n] > arr[r]{
		n = r
	}

	if arr[p] < arr[n]{
		return
	}

	arr[p], arr[n] = arr[n], arr[p]

	heap(arr, n, end)
}

func HeapSort2(arr []int)  {
	if len(arr) <= 1{
		return
	}
	l := len(arr)
	for i := l / 2; i >= 0; i--{
		heap(arr, i, l - 1)
	}

	for i := l - 1; i >= 0; i--{
		arr[0], arr[i] = arr[i], arr[0]
		heap(arr, 0, i - 1)
	}
}

func heapfiy3(arr []int, p, end int)  {
	l := p * 2 + 1
	if l > end{
		return
	}
	r := p * 2 + 2
	n := l
	if r <= end && arr[n] > arr[r]{
		n = r
	}

	if arr[p] < arr[n]{
		return
	}

	arr[p], arr[n] = arr[n], arr[p]
	heapfiy3(arr, n, end)
}

func heap3(arr []int)  {
	len := len(arr)
	m := len / 2
	for i := m; i >= 0; i--{
		heapfiy3(arr, i, len-1)
	}
	for i := len - 1; i >= 0; i--{
		arr[0], arr[i] = arr[i], arr[0]
		heapfiy(arr, 0, i)
	}
}