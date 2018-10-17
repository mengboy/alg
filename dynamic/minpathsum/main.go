package main

import "math"

func minpathsum(grid [][]int) int{
	r := len(grid)
	c := len(grid[0])
	if r == 0 && c == 0{
		return 0
	}

	ret := IntMatrix(r, c)
	ret[0][0] = grid[0][0]

	for i := 1; i < r; i++{
		ret[i][0] = ret[i - 1][0] + grid[i][0]
	}
	for i := 1; i < c; i++{
		ret[0][i] = ret[0][i - 1] + grid[0][i]
	}

	for i := 1; i < r; i++{
		for j := 1; j < c; j++{
			ret[i][j] = grid[i][j] + int(math.Min(float64(ret[i - 1][j]), float64(ret[i][j - 1])))
		}
	}
	return ret[r - 1][c - 1]

}

func recur(grid [][]int, x, y int) int  {
	if len(grid) == 0 && len(grid[0]) == 0{
		return 0
	}

	if x == 0{
		return recur(grid, x, y - 1) + grid[x][y]
	}

	if y == 0{
		return recur(grid, x - 1, y) + grid[x][y]
	}

	return grid[x][y] + int(math.Min(float64(recur(grid, x - 1, y)), float64(recur(grid, x, y - 1))))
}

func main()  {
	
}

func IntMatrix(rows, cols int) [][]int {
	m := make([][]int, rows)
	for i := 0; i < rows; i++ {
		m[i] = make([]int, cols)
	}
	return m
}
