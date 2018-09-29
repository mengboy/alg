package main

import (
	"fmt"
)

func main() {
	//arr := []int{
	//	4,
	//	5,
	//	6,
	//	7,
	//	0,
	//	1,
	//	2,
	//}

	//arr := []int{
	//	0,
	//	1,
	//	2,
	//	4,
	//}

	//fmt.Println(bianrySearchValue(arr, 0, 4, 6))

	arr := []int{
		0,
		1,
		3,
	}

	//fmt.Println(search(arr, 1))

	fmt.Println(bianrySearchValue(arr, 0, 3, 0))
}

func search(nums []int, target int) int {
	if len(nums) == 0 {
		return -1
	}

	if len(nums) == 1 {
		if target == nums[0] {
			return 0
		} else {
			return -1
		}
	}

	pointIndex := bianrySearch(nums, 0, len(nums))

	if target == nums[0] {
		return 0
	}

	// 递增序列
	if pointIndex == len(nums) {
		return bianrySearchValue(nums, 0, pointIndex, target)
	}

	if target > nums[0] {
		return bianrySearchValue(nums, 0, pointIndex, target)
	} else {
		return bianrySearchValue(nums, pointIndex, len(nums), target)
	}

}

/**
*  获取旋转位置
 */
func bianrySearch(nums []int, start, end int) int {
	t := nums[0]
	mid := (end-start)/2 + start

	for start < end {
		mid = (end-start)/2 + start
		if nums[mid] > t {
			start = mid
			if start+1 == end && end == len(nums) {
				return end
			}
		} else {
			if nums[mid] < t && nums[mid-1] >= t {
				return mid
			}
			if nums[mid] < t && nums[mid-1] < t {
				end = mid
			}
		}
	}
	return mid
}

func bianrySearchValue(nums []int, start, end, target int) int {

	if target < nums[start] || target > nums[end-1] {
		return -1
	}

	for start <= end {
		mid := (end-start)/2 + start
		if nums[mid] < target {
			start = mid + 1
		} else if nums[mid] > target {
			end = mid - 1
		} else {
			return mid
		}
	}
	return -1
}
