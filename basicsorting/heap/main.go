package main

func main()  {

}

func HeapSort(arr []int)  {
	if len(arr) <= 1{
		return
	}
	l := len(arr)
	for i := len(arr) / 2; i >= 0; i--{
		heapfiy(arr, i, l - 1)
	}
	for i := l - 1; i >= 0; i++{
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
