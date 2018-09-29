package main

import (
	"fmt"
	"sort"
)

func main() {
	nums := []int{
		1, 0, -1, 0, -2, 2,
	}

	fmt.Println(fourSum(nums, 0))
}

func fourSum(nums []int, target int) [][]int {

	if len(nums) < 4 {
		return nil
	}

	if len(nums) == 4 {
		if nums[0]+nums[1]+nums[2]+nums[3] == target {
			return [][]int{
				nums,
			}
		}
	}

	fmt.Println(nums)

	arrs := make([][]int, 0)
	for i := 0; i < len(nums); i++ {
		t := target - nums[i]
		leftNum := make([]int, 0)
		leftNum = append(leftNum, nums[:i]...)
		leftNum = append(leftNum, nums[i+1:]...)
		tarrs := threeSum(leftNum, t)
		fmt.Println(tarrs)
		for _, tarr := range tarrs {
			tarr := append(tarr, nums[i])
			sort.Ints(tarr)
			if !contain(arrs, tarr) {
				arrs = append(arrs, tarr)
			}
		}
	}

	return arrs
}

func threeSum(nums []int, target int) [][]int {
	sort.Ints(nums)
	arrs := make([][]int, 0)
	for i := 0; i < len(nums); i++ {
		//fmt.Println(nums)
		l := i + 1
		r := len(nums) - 1
		for l < r {
			sum := nums[i] + nums[l] + nums[r]
			if sum == target {
				tarr := []int{
					nums[i],
					nums[l],
					nums[r],
				}
				if !contain(arrs, tarr) {
					arrs = append(arrs, tarr)
				}
			}
			t := sum - target
			if t < 0 {
				l++
			} else {
				r--
			}
		}
	}
	return arrs
}

func contain(arrs [][]int, arr []int) bool {
	for _, v := range arrs {
		if len(v) == 3 {
			if v[0] == arr[0] && v[1] == arr[1] && v[2] == arr[2] {
				return true
			}
		}
		if len(v) == 4 {
			if v[0] == arr[0] && v[1] == arr[1] && v[2] == arr[2] && v[3] == arr[3] {
				return true
			}
		}
	}
	return false
}
