package main

import "fmt"

func main() {
	fmt.Println(searchRange([]int{
		1,
		1,
		2,
	}, 1))
}

func searchRange(nums []int, target int) []int {
	res := []int{
		-1,
		-1,
	}
	if len(nums) == 0 {
		return res
	}
	i := binarySearch(nums, 0, len(nums)-1, target)
	if i == -1 {
		return res
	}

	l := i
	for ; l >= 0; l-- {
		if nums[l] == target {
			continue
		} else {
			break
		}
	}
	l++

	r := i

	for ; r < len(nums); r++ {
		if nums[r] == target {
			continue
		} else {
			break
		}
	}
	r--

	return []int{
		l,
		r,
	}
}

func binarySearch(nums []int, start, end, target int) int {

	if target < nums[start] || target > nums[end] {
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
