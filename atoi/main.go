package main

import (
	"fmt"
	"math"
	"strconv"
	"strings"
)

/**
* 需要注意首字符：'+'
* 多个'+'、'-'连续
* 非首位的'+'、'-'
 */

func main() {
	fmt.Println(myAtoi("-2147483649"))
}

func myAtoi(str string) int {
	if len(str) == 0 {
		return 0
	}
	max := math.MaxInt32
	min := math.MinInt32
	numStr := "1234567890"
	str = strings.TrimSpace(str)
	result := ""
	//myint := 0
	for i, c := range str {
		if t := strings.Contains(numStr, string(c)); t == false {
			if c == '-' {
				if i > 0 {
					break
				}
				result += string(c)
				continue
			}
			if c == '+' {
				if i > 0 {
					break
				}
				continue
			}
			break
		} else {
			result += string(c)
			//myint = myint * 10 + int(c - '0')
			//fmt.Println(myint)
		}
	}

	//if result[0] == '-'{
	//	fmt.Println(0 - myint)
	//}

	if len(result) == 0 {
		return 0
	}

	n, _ := strconv.ParseInt(result, 10, 64)

	if n > int64(max) {
		return max
	}
	if n < int64(min) {
		return min
	}

	return int(n)
}

// 不使用库函数
// myint := 0
//for i := 0; teststr[i] ; i++
//{
//myint = myint*10+teststr[i]-'0';
//}
