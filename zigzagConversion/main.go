package main

import "fmt"

/**
z字形变换
*/

/**
* 需要注意 numRows 为1 的情况
 */

func main() {

	s := "A"

	fmt.Println(convert(s, 1))
}

/**
n
PAYPALISHIRING
3
PAHNAPLSIIGYIR
PAHNAPLSIIGYIR

4
PINALSIGYAHRPI
PINALSIGYAHRPI


*/
func convert(s string, numRows int) string {

	if numRows == 1 {
		return s
	}

	cs := ""
	step := numRows*2 - 2
	for i := 0; i < numRows; i++ {
		if i == 0 || i == numRows-1 {
			for j := i; j < len(s); j += step {
				cs += string(s[j])
			}
		} else {
			for j := i; j < len(s); j += step {
				cs += string(s[j])
				// 第n个循环
				t := ((j-i)/step + 1) * numRows
				// 第n个
				t += ((j - i) / step) * (numRows - 2)
				x := t + (numRows - i) - 1 - 1
				if x < len(s) {
					cs += string(s[x])
				}

			}
		}
	}
	return cs
}
