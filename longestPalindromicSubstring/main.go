package main

import "fmt"

/**
* 最长回文子串
* 需要注意连续相等字符的情况
 */

func main() {
	s := "babad"
	//t := "ac"
	fmt.Println(longestPalindrome(s))
}

func longestPalindrome(s string) string {
	l := len(s)
	if l < 2 {
		return s
	}

	subStr := ""
	for i := 1; i < l; i++ {
		tstr := sub(i, s)
		if len(subStr) < len(tstr) {
			subStr = tstr
		}
	}
	return subStr
}

func sub(mid int, s string) string {
	left := mid - 1
	right := mid + 1

	for left >= 0 || right < len(s) {
		if left >= 0 && s[mid] == s[left] {
			left--
			continue
		}
		if right < len(s) && s[mid] == s[right] {
			right++
			continue
		}
		break
	}

	for left >= 0 && right < len(s) {
		if s[left] == s[right] {
			left--
			right++
		} else {
			break
		}
	}
	return s[left+1 : right]
}
