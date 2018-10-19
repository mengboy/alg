package main

import "fmt"

func main()  {
	arr := []int{
		20, 1, 2, 3, 7, 4, 9, 10, 19, 17, 15,
	}
	BubbleSort(arr)
}

func BubbleSort(arr []int)  {
	l := len(arr)
	for i := 0; i < l; i++{
		for j := 0; j < l - i - 1; j++{
			if arr[j] > arr[j + 1]{
				arr[j], arr[j + 1] = arr[j + 1], arr[j]
			}
		}
		fmt.Println(arr)
	}

}
