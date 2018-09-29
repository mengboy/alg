package main

import (
	"fmt"
	"math"
)

func main() {
	//nums := []int{
	//	-2,1,-3,4,-1,2,1,-5,4,
	//}

	nums := []int{
		-2,
	}
	fmt.Println(maxSubArray(nums))
}

func maxSubArray(nums []int) int {
	if len(nums) == 0 || nums == nil {
		return 0
	}
	sum := 0
	maxSum := math.MinInt32
	for _, v := range nums {
		if sum < 0 {
			sum = 0
		}
		sum += v
		if maxSum < sum {
			maxSum = sum
		}

	}
	return maxSum
}

//   -3 | 4 | -1 | 2

func mergeMaxSubArray(nums []int) int {
	return 0
}
