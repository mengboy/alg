package main

import (
	"fmt"
	"math"
)

/**
* 背包问题
 */

func main()  {
	w1 := []int{
		2, 1, 3, 2,
	}
	v1 := []int{
		3, 2, 4, 2,
	}
	W1 := 5
	w2 := []int{
		6, 3, 4,
	}
	v2 := []int{
		23, 13, 11,
	}
	W2 := 17
	//var itemTake [len(w1) + 1]bool
	bp1,  itemTake := backpack(W1, w1, v1)
	fmt.Println("max v:", bp1)
	for i, v := range itemTake{
		if v {
			fmt.Println("item ", i, " w:", w1[i], ", value: ", v1[i])
		}
	}

	bp2, itemTake2 := backpack(W2, w2, v2)
	fmt.Println("max v", bp2)
	for i, v := range itemTake2{
		if v {
			fmt.Println("item ", i, " w:", w2[i], ", value: ", v2[i])
		}
	}

	bp3 := repeatBackPack(W2, w2, v2)
	fmt.Println("max v", bp3)

}

func IntMatrix(rows, cols int) [][]int {
	m := make([][]int, rows)
	for i := 0; i < rows; i++ {
		m[i] = make([]int, cols)
	}
	return m
}

func  BoolMatrix(rows, cols int) [][]bool {
	m := make([][]bool, rows)
	for i := 0; i < rows; i++ {
		m[i] = make([]bool, cols)
	}
	return m
}

// 不重复取
func backpack(W int, w[]int, v []int) (int, []bool) {
	N := len(w)

	dp := IntMatrix(N + 1, W + 1)

	matrix := BoolMatrix(N + 1, W + 1)

	itemTake := make([]bool, N + 1)

	for i := 0; i < N; i++{
		for j := 0; j <= W; j++{
			// 当前元素重量大于背包容纳量
			if (w[i] > j){
				// 设置前i种商品最大值和前i-1种商品最大值一样
				dp[i + 1][j] = dp[i][j]
			}else {
				//
				dp[i + 1][j] = int(math.Max(float64(dp[i][j]), float64(dp[i][j - w[i]] + v[i])))
				matrix[i][j] = (dp[i][j - w[i]] + v[i] > dp[i][j])
			}
		}
	}


	var j int
	j = W
	for i := N - 1; i >= 0; i-- {
		if matrix[i][j]{
			itemTake[i] = true
			j -= w[i]
		}else {
			itemTake[i] = false
		}
	}

	return dp[N][W], itemTake
}

// 可以重复取
func repeatBackPack(W int, w []int, v []int) int {
	N := len(w)
	dp := IntMatrix(N + 1, W + 1)
	for i := 0; i < N; i++{
		for j := 0; j <= W; j++{
			if (w[i] > j){
				// 初始化一些初始值
				dp[i + 1][j] = dp[i][j]
			}else {
				// 
				dp[i + 1][j] = int(math.Max(float64(dp[i][j]), float64(dp[i + 1][j - w[i]] + v[i])))
				fmt.Println("i ", i, "w", w[i], "v", v[i])
			}
		}
	}
	return dp[N][W]
}

// lintcode 92
func solution(m int, w []int) int{
	dp := IntMatrix(len(w) + 1, m + 1)
	for i := 0; i < len(w); i++{
		for j := 0; j <= m; j++{
			if w[i] > j{
				dp[i + 1][j] = dp[i][j]
			}else {
				dp[i + 1][j] = int(math.Max(float64(dp[i][j]), float64(dp[i][j - w[i]] + w[i])))
			}
		}
	}
	return dp[len(w) + 1][m]
}

