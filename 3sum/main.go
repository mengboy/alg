package main

import (
	"fmt"
	"sort"
)

/**
* 三个数和为0
* 有参考
* 固定一个数,通过移动两个指针找第三个数
 */

func main() {
	nums := []int{
		//  -1, 0, 1, 2, -1, -4,
		//  0, 1, 1, 1,
		//  -1, 0,  1,
		-2, 0, 1, 1, 2,
	}

	//nums := []int{
	//	//0, 1, 1,
	//	-1, -1, 0,
	//}

	fmt.Println(threeSum2(nums))
}

func threeSum(nums []int) [][]int {

	if len(nums) < 3 {
		return nil
	}

	if len(nums) == 3 {
		sum := nums[0] + nums[1] + nums[2]
		if sum == 0 {
			return [][]int{
				nums,
			}
		} else {
			return nil
		}
	}

	sort.Ints(nums)

	if nums[0] > 0 || nums[len(nums)-1] < 0 {
		return nil
	}

	start := 0
	end := len(nums) - 1

	arrs := make([][]int, 0)
	for start < end {
		n := 0
		t2 := nums[start] + nums[end]
		if t2 > 0 {
			n = start + 1
			for n < end {
				t3 := t2 + nums[n]
				if t3 > 0 {
					end--
					break
				}

				if t3 == 0 {
					arr := []int{
						nums[start],
						nums[n],
						nums[end],
					}
					arrs = append(arrs, arr)
					start++
					end--
					break
				}

				n++
			}
			if n == end {
				break
			}
		}

		if t2 < 0 {
			n = end - 1
			for n > start {
				t3 := t2 + nums[n]
				if t3 < 0 {
					start++
					break
				}

				if t3 == 0 {
					arr := []int{
						nums[start],
						nums[n],
						nums[end],
					}
					arrs = append(arrs, arr)
					start++
					end--
					break
				}
				n--
			}
			if n == start {
				break
			}
		}

		if t2 == 0 {
			n = start + 1
			for n < end {
				t3 := t2 + nums[n]
				if t3 > 0 {
					end--
					break
				}
				if t3 == 0 {
					arr := []int{
						nums[start],
						nums[n],
						nums[end],
					}
					arrs = append(arrs, arr)
					start++
					end--
					break
				}
				n++
			}

			if n == end {
				break
			}

		}

	}
	return arrs
}

func threeSum2(nums []int) [][]int {
	if len(nums) < 3 {
		return nil
	}

	if len(nums) == 3 {
		sum := nums[0] + nums[1] + nums[2]
		if sum == 0 {
			return [][]int{
				nums,
			}
		} else {
			return nil
		}
	}
	sort.Ints(nums)

	arrs := make([][]int, 0)
	for i, _ := range nums {
		l := i + 1
		r := len(nums) - 1
		for l < r {
			t := nums[i] + nums[l] + nums[r]
			if t == 0 {
				arr := []int{
					nums[i],
					nums[l],
					nums[r],
				}
				if !contain(arrs, arr) {
					arrs = append(arrs, arr)
				}
			}
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
		if v[0] == arr[0] && v[1] == arr[1] && v[2] == arr[2] {
			return true
		}
	}
	return false
}
