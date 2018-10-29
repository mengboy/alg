package main

import (
	"fmt"
	"math"
)


/**
* 两数相除
* 29
 */

/**
* 需要优化
 */

func main() {

	//2
	//2147483647

	fmt.Println(math.MaxInt32)

	fmt.Println(divide(-1, 1))
	//fmt.Println(divide(10, 3))
}

func divide(dividend int, divisor int) int {

	if dividend == math.MinInt32 && divisor == -1 {
		return math.MaxInt32
	}

	if divisor == 1 {
		return dividend
	}

	if dividend == -1 {
		return -dividend
	}

	flag := false

	if (divisor < 0 && dividend < 0) || (dividend > 0 && divisor > 0) {
		flag = true
	}

	if divisor > 0 {
		divisor = -divisor
	}

	if dividend > 0 {
		dividend = -dividend
	}

	if dividend == 0 || divisor < dividend {
		return 0
	}

	if dividend == divisor {
		if flag {
			return 1
		} else {
			return -1
		}
	}

	t := divisor

	i := 0
	//for dividend - divisor > 0{
	//	divisor = divisor << 1
	//	i++
	//}
	//divisor = divisor >> 1
	//i--
	//
	//i += i

	for dividend-divisor <= 0 {
		divisor += t
		i++
	}
	//i--

	if flag {
		return i
	} else {
		return -i
	}
}

func divide2(dividend int, divisor int) int{
	sub := dividend
	res := 0

	for sub >= divisor{
		tmp := divisor
		i := 1
		for sub > tmp  {
			sub = sub - tmp
			res += i
			// 通过移位增加位数
			i = i << 1
			tmp = tmp << 1
		}
	}
	return res
}