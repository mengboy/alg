package main

import "fmt"

func main() {
	fmt.Println(climbStairs(45))
}

/**
* 1 1
* 2 2
* 3 3
* 4       1111  112  13  121  211 31
 */

func climbStairs(n int) int {
	mark := map[int]int{}
	return fbi(n, mark)

}

func fbi(n int, mark map[int]int) int {
	if n == 1 {
		return 1
	}

	if n == 2 {
		return 2
	}

	if mark[n] > 0 {
		return mark[n]
	}

	ret := fbi(n-1, mark) + fbi(n-2, mark)

	mark[n] = ret

	return ret
}
