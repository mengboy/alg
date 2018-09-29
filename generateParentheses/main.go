package main

import "fmt"

/**
* 括号生成
* 参考
 */

func main() {
	fmt.Println(generateParenthesis(3))
}

func generateParenthesis(n int) []string {
	if n == 0 {
		return nil
	}

	strs := make([]string, 0)

	dfs(0, 0, "", &strs, n)

	return strs
}

func dfs(left, right int, buff string, strs *[]string, n int) {
	if left == n && right == n {
		*strs = append(*strs, buff)
	}
	if left < n {
		dfs(left+1, right, buff+"(", strs, n)
	}
	if left > right {
		dfs(left, right+1, buff+")", strs, n)
	}
}
