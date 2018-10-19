package main

import (
	"math"
)

func main()  {

}

func SelectSort(arr []int)  {
	l := len(arr)
	for i := 0; i < l; i++{
		max := math.MaxInt32
		maxi := math.MaxInt32
		for j := 0; j < l - i; j++{
			if arr[j] > max{
				max = arr[j]
				maxi = j
			}
		}
		arr[maxi], arr[l - 1] = arr[l - 1], arr[maxi]
	}
}
