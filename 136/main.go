package main

import "fmt"

/**
* 只出现一次对数字
* 数组中只有一个数字出现一次，其他均出现两次
 */

func main()  {
	//arr := []int{
	//	2,2,1,
	//}
	arr2 := []int{
		4,1,2,1,2,
	}
	fmt.Println(singleNumber(arr2))

	fmt.Println(2 | 2, 2 & 2, 2 ^ 2, 0 ^ 3 )
}


func singleNumber(nums []int) int {
	var t = nums[0]
	for i := 0; i < len(nums) - 1; i++ {
		t = t ^ nums[i + 1]
	}
	return t
}