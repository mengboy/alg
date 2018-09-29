package main

import (
	"fmt"
	"math"
	"sort"
)

func main() {

	nums := []int{
		0, 2, 1, -3,
	}

	fmt.Print(threeSumClosest(nums, 1))
}

func threeSumClosest(nums []int, target int) int {
	if len(nums) == 0 {
		return 0
	}
	sum := 0
	if len(nums) < 4 {
		for _, num := range nums {
			sum += num
		}
		return sum
	}

	sort.Ints(nums)
	diff := math.MaxFloat64
	for i, _ := range nums {
		l := i + 1
		r := len(nums) - 1
		for l < r {
			t := nums[i] + nums[l] + nums[r]
			realDiff := float64(t - target)
			if realDiff == 0 {
				sum = t
				return sum
			}
			realDiffAbs := math.Abs(realDiff)
			if realDiffAbs < diff {
				diff = realDiffAbs
				sum = t
			}

			if realDiff < 0 {
				l++
			} else {
				r--
			}
		}
	}

	return sum
}
