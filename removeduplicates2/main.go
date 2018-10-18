package main

import "fmt"

func main()  {
	arr := []int{
		1,1,1,2,2,3,
	}
	arr2 := []int{
		0,0,1,1,1,1,2,3,3,
	}

	fmt.Println(removeDuplicates(arr))
	fmt.Println(removeDuplicates2(arr2))

}


func removeDuplicates(nums []int) int {
	l := len(nums)
	if l == 0 || l == 1{
		if l == 1{
			fmt.Println(nums[0])
		}
		return l
	}

	// 拿当前值和前两个值相比较，出现任意一个不相等，更新索引，更新值，若相等不做任何操作
	i := 1
	for j := 2; j < l; j++{
		if nums[j] != nums[i] || nums[j] != nums[i - 1]{
			i++
			nums[i] = nums[j]
		}
	}

	for j := 0; j <= i; j++{
		fmt.Print(nums[j], " ")
	}
	fmt.Println()
	return i + 1
}


func removeDuplicates2(nums []int) int {
	if len(nums) < 3 {
		return len(nums)
	}
	size := 0
	rep := 0 // 重复个数
	for i := 1; i != len(nums); i++ {
		if nums[i] == nums[i-1] {
			rep++
			if rep > 1 {
				// 重复两个以上，需要更新
				size++
			}
		} else {
			rep = 0
		}
		if size > 0 {
			nums[i-size] = nums[i]
		}
	}
	l := len(nums) - size
	for j := 0; j < l; j++{
		fmt.Print(nums[j], " ")
	}
	fmt.Println()
	return  l
}