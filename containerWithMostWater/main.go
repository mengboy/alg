package main

import (
	"fmt"
)

/**
* 注意 高度相等的情况
 */

func main() {
	height := []int{
		1, 8, 6, 2, 5, 4, 8, 3, 7,
	}
	fmt.Println(maxArea(height))
}

func maxArea(height []int) int {

	if len(height) < 2 {
		return 0
	}
	area := 0
	i := 0
	j := len(height) - 1
	for i < j {
		a := 0
		if height[i] < height[j] {
			a = height[i] * (j - i)
			i++
		} else {
			a = height[j] * (j - i)
			j--
		}
		if area < a {
			area = a
		}
	}
	return area
}
