package main

import (
	"fmt"
)

/**
* 删除数组重复项
 */

func main()  {
	arr := []int{
		1, 1, 2, 2, 3, 3, 4, 4, 4, 4, 5, 5, 5,
	}
	fmt.Println(removeDuplicates(arr))
}


func removeDuplicates(nums []int) int {
	l := len(nums)

	if l == 0 || l == 1{
		if l == 1{
			fmt.Println(nums[0])
		}
		return l
	}

	i := 1
	for j := 0; j < l - 1; j++{
		if nums[j] == nums[j + 1]{
			continue
		}else {
			nums[i] = nums[j + 1]
			i++
	 	}
	}

	for j := 0; j < i - 1; j++{
		fmt.Println(nums[j])
	}
	return i
}