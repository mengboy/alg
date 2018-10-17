package main

import (
	"fmt"
	"math"
)

/**
* 计算三角形最小路径和
 */

func main()  {
	t := [][]int{
		{2},
		{3, 4},
		{6, 5, 4},
		{4, 1, 8, 3},
		{8, 5, 4, 9, 1},
	}

	res := dfs(0, 0, 0, math.MaxInt32, t)
	fmt.Println(res)
	fmt.Println(int(dfs2(0, 0, t)))
}

func dfs(x, y int, sum int, result int, triangle [][]int) int {
	n := len(triangle)
	if x == n {
		if sum < result{
			result = sum
		}
		return result
	}

	if y < len(triangle[x]){
		sum += triangle[x][y]
	}

	result = dfs(x + 1, y, sum, result, triangle)

	result = dfs(x + 1, y + 1, sum, result, triangle)

	return result
}

func dfs2(x, y int, triangle [][]int) float64  {
	if x == len(triangle){
		return 0
	}
	if y >= len(triangle[x]){
		return 0
	}
	return math.Min(dfs2(x + 1, y, triangle), dfs2(x + 1, y + 1, triangle)) + float64(triangle[x][y])
}
