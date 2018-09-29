package main

import "fmt"

func main() {
	fmt.Println(letterCombinations("234"))
}

func letterCombinations(digits string) []string {
	mapnum2letter := map[string]string{
		"2": "abc",
		"3": "def",
		"4": "ghi",
		"5": "jkl",
		"6": "mno",
		"7": "pqrs",
		"8": "tuv",
		"9": "wxyz",
	}

	if len(digits) == 0 {
		return nil
	}

	result := make([]string, 0)
	letterStr := make([]string, 0)

	for _, digit := range digits {
		letterStr = append(letterStr, mapnum2letter[string(digit)])
	}

	for _, v := range letterStr[0] {
		result = append(result, string(v))
	}

	for i := 1; i < len(letterStr); i++ {
		result = fci(result, letterStr[i])
	}

	return result
}

func fci(strs []string, str string) []string {

	r := make([]string, 0)

	for _, c := range str {
		for _, s := range strs {
			ts := s + string(c)
			r = append(r, ts)
		}
	}

	return r
}
