package main

import "fmt"

func main()  {
	arr := []int{
		20, 1, 2, 3, 7, 4, 9, 10, 19, 17, 15,
	}

	fmt.Println(MergeSort(arr))
}

func MergeSort(arr []int) []int {
	if len(arr) == 0 || len(arr) == 1{
		return arr
	}
	mid := len(arr) / 2
	left := MergeSort(arr[0:mid])
	right := MergeSort(arr[mid:len(arr)])
	return merge(left, right)
}

func merge(arr1 []int, arr2 []int) []int {
	l1 := len(arr1)
	l2 := len(arr2)
	newArr := make([]int, 0)
	i,j := 0,0
	for i < len(arr1) && j < len(arr2){
		if arr1[i] < arr2[j]{
			newArr = append(newArr, arr1[i])
			i++
		}else {
			newArr = append(newArr, arr2[j])
			j++
		}
	}
	if i < l1{
		newArr = append(newArr, arr1[i:]...)
		}

	if j < l2 {
		newArr = append(newArr, arr2[j:]...)
		}

	return newArr
}

func merge2(arr1 []int, arr2 []int)  []int {
	l1 := len(arr1)
	l2 := len(arr2)
	newArr := make([]int, 0)
	i, j := 0, 0
	for ; i < l1 && j < l2;{
		if arr1[i] < arr2[j]{
			newArr = append(newArr, arr1[i])
			i++
		}else {
			newArr = append(newArr, arr2[j])
			j--
		}
	}
	if i < l1{
		newArr = append(newArr, arr1[i:]...)
	}
	if j < l2{
		newArr = append(newArr, arr2[j:]...)
	}

	return newArr
}

func mergeSort(arr []int) []int  {
	if len(arr) <= 1{
		return arr
	}
	m := len(arr) / 2
	left := arr[0:m]
	right := arr[m:len(arr)]
	return merge2(left, right)
}