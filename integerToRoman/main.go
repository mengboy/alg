package main

import (
	"fmt"
	"strconv"
)

/**
* 整数转罗马数字
 */

func main() {
	fmt.Print(intToRoman(1994))
}

func intToRoman(num int) string {
	roman := ""
	strNum := strconv.FormatInt(int64(num), 10)
	if len(strNum) == 4 {
		roman += Thousand2Roman(int(strNum[0] - '0'))
		roman += Hundreds2Roman(int(strNum[1] - '0'))
		roman += Ten2Roman(int(strNum[2] - '0'))
		roman += One2Roman(int(strNum[3] - '0'))
		return roman

	}

	if len(strNum) == 3 {
		roman += Hundreds2Roman(int(strNum[0] - '0'))
		roman += Ten2Roman(int(strNum[1] - '0'))
		roman += One2Roman(int(strNum[2] - '0'))
		return roman
	}

	if len(strNum) == 2 {
		roman += Ten2Roman(int(strNum[0] - '0'))
		roman += One2Roman(int(strNum[1] - '0'))
		return roman
	}

	if len(strNum) == 1 {
		roman += One2Roman(int(strNum[0] - '0'))
		return roman
	}
	return roman
}

func Thousand2Roman(thousand int) string {
	res := ""
	for i := 0; i < thousand; i++ {
		res += "M"
	}
	return res
}

func Hundreds2Roman(hundreds int) (res string) {

	if hundreds == 9 {
		res = "CM"
		return
	}

	if hundreds == 4 {
		res = "CD"
		return
	}

	if hundreds < 4 {
		for i := 0; i < hundreds; i++ {
			res += "C"
		}
		return
	}

	if hundreds == 5 {
		res = "D"
		return
	}

	if hundreds > 5 {
		res += "D"
		for i := 0; i < (hundreds - 5); i++ {
			res += "C"
		}
		return
	}

	return
}

func Ten2Roman(ten int) (res string) {
	if ten == 9 {
		res = "XC"
		return
	}

	if ten == 4 {
		res = "XL"
		return
	}

	if ten < 4 {
		for i := 0; i < ten; i++ {
			res += "X"
		}
		return
	}

	if ten == 5 {
		res = "L"
		return
	}

	if ten > 5 {
		res += "L"
		for i := 0; i < (ten - 5); i++ {
			res += "X"
		}
		return
	}
	return
}

func One2Roman(one int) (res string) {

	if one == 9 {
		res = "IX"
		return
	}

	if one == 4 {
		res = "IV"
		return
	}

	if one < 4 {
		for i := 0; i < one; i++ {
			res += "I"
		}
		return
	}

	if one == 5 {
		res = "V"
		return
	}

	if one > 5 {
		res += "V"
		for i := 0; i < (one - 5); i++ {
			res += "I"
		}
		return
	}
	return

}
