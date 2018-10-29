package main

import (
	"fmt"
	"sort"
)

/**
* 求数组中两个数和为指定数
 */
func main()  {
	arr := []int{
		2, 3, 4, 4, 6, 8,
	}
	//i1, i2 := result2(arr, 8)

	//fmt.Println(i1, i2)
}

// x1 + x2 = target
// x2 = target - x1
// 使用map保存数组元素和索引
// 拿空间换时间

func result(arr []int, target int) []int  {
	_map := make(map[int]int)
	for i, v := range arr{
		if index, ok := _map[target - v]; ok{
			return []int{i, index}
		}else {
			_map[v] = i
		}
	}
	return []int{-1, -1}
}


// 解2
// 拿时间换空间
// 对原素组排序，排序后使用双指针

func result2(arr []int, target int)  []int {
	sort.Ints(arr)
	for i, j := 0, len(arr) - 1; i < j; {
		if arr[i] + arr[j] == target{
			return []int{
				i, j,
			}
		}else {
			if arr[i] + arr[j] > target{
				j--
			}else {
				i++
			}
		}
	}
	return []int{
		-1, -1,
	}
}


