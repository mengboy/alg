package heap

import (
	"fmt"
)

type Heap struct {
	Data []int
	Len int
	Min bool
}

func NewMin() *Heap {
	data := make([]int, 0)
	return &Heap{
		Data:data,
		Len:0,
		Min:true,
	}
}

func (h *Heap) Insert(x int)  {
	h.Data = append(h.Data, x)
	h.Len++
	h.ShiftUp()
}

func (h *Heap) Pop() *int  {
	if h.Len == 0{
		return nil
	}
	f := h.Data[0]
	last := h.Data[h.Len - 1]
	if h.Len == 1{
		h.Data = h.Data[0:0]
		h.Len--
		return &f
	}
	h.Data = append([]int{last}, h.Data[1:h.Len - 1]...)
	h.Len--
	h.ShiftDown()
	return &f
}

func (h *Heap) ShiftUp()  {
	if h.Len == 0 || h.Len == 1{
		return
	}
	for i, parent := h.Len - 1, h.Len - 1; i > 0; i = parent{
		parent = i >> 1
		// 子节点小于父节点
		if h.Less(h.Data[i], h.Data[parent]){
			h.Data[parent], h.Data[i] = h.Data[i], h.Data[parent]
		}else {
			break
		}

	}
	fmt.Println(h.Data)
}

func (h *Heap) ShiftDown()  {
	if h.Len == 0 || h.Len == 1{
		return
	}

	for i, child := 0, 1; i <  h.Len && i << 1 + 1 < h.Len; i = child{
		child = i << 1 + 1
		if child+1 <= h.Len -1 && h.Less(h.Data[child+1], h.Data[child]) {
			child++
		}

		if h.Less(h.Data[i], h.Data[child]) {
			break
		}

		h.Data[i], h.Data[child] = h.Data[child], h.Data[i]
	}
}

func (h *Heap) Heapfiy()  {
	if h.Len == 0 || h.Len == 1{
		return
	}
	for i := h.Len / 2; i > -1; i--{
		heap(h.Data, i, h.Len - 1)
	}
	fmt.Println(h.Data)
	for i := h.Len-1; i > 0; i-- {
		h.Data[i], h.Data[0] = h.Data[0], h.Data[i]
		heapfiy(h.Data, 0, i-1)
	}
	//fmt.Println(h.Data)
}

func heap(data []int, i, e int)  {
	l := i << 1 + 1
	if l > e {
		return
	}
	r := i << 1 + 2
	n := l
	if r > e{
		return
	}
	if r <= e && data[r] > data[l]{
		n = r
	}

	if data[i] > data[n]{
		return
	}
	data[i], data[n] = data[n], data[i]
	heap(data, n, e)
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

func (h *Heap) Less(x, y int) bool {
	if h.Min{
		return x < y
	}else {
		return x > y
	}
}

