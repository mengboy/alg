package main

import (
	"fmt"
)

/**
* 最长子串的长度
* abcabcbb 3
 */

func main() {
	s := " bbcaa"
	//s := "abba"
	result := 0
	calLen(&result, s)
	fmt.Println(result)
}

func lengthOfLongestSubstring(s string) int {
	result := 0
	calLen(&result, s)
	return result
}

func calLen(result *int, s string) {
	subLen := 0
	charNum := make(map[uint8]int)
	for i := 0; i < len(s); i++ {
		if charNum[s[i]] > 0 {
			newS := s[charNum[s[i]]:len(s)]
			calLen(result, newS)
			break
		}
		subLen++
		charNum[s[i]] = i + 1
		if *result < subLen {
			*result = subLen
		}
	}
}

func step(start int, s string) int {
	charNum := make(map[uint8]int)
	result := 0
	subLen := 0
	step := start
	for i := start; i < len(s); i++ {
		if charNum[s[i]] > 0 {
			step = charNum[s[i]] - 1
			subLen = i - step
			if result < subLen {
				result = subLen
			}
			// 更新最近的一个重复字符位置
			charNum[s[i]] = i + 1
			continue
		}
		subLen++
		charNum[s[i]] = i + 1
		if result < subLen {
			result = subLen
		}
	}
	return result
}
