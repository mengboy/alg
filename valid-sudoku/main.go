package main

import "fmt"

func main() {
	fmt.Println(byte('.'))
	fmt.Println(byte('0'))
	fmt.Println(byte('1'))
	fmt.Println(byte('2'))
	fmt.Println(byte('3'))
	fmt.Println(byte('4'))
	fmt.Println(byte('5'))
	fmt.Println(byte('6'))
	fmt.Println(byte('7'))
	fmt.Println(byte('8'))
	fmt.Println(byte('9'))

	board := [][]byte{
		{'8', '3', '.', '.', '7', '.', '.', '.', '.'},
		{'6', '.', '.', '1', '9', '5', '.', '.', '.'},
		{'.', '9', '8', '.', '.', '.', '.', '6', '.'},
		{'8', '.', '.', '.', '6', '.', '.', '.', '3'},
		{'4', '.', '.', '8', '.', '3', '.', '.', '1'},
		{'7', '.', '.', '.', '2', '.', '.', '.', '6'},
		{'.', '6', '.', '.', '.', '.', '2', '8', '.'},
		{'.', '.', '.', '4', '1', '9', '.', '.', '5'},
		{'.', '.', '.', '.', '8', '.', '.', '7', '9'},
	}
	fmt.Println(line(board))
	fmt.Println(row(board))
	fmt.Println(ninexnine(board))
}

func isValidSudoku(board [][]byte) bool {

	return line(board) && row(board) && ninexnine(board)

}

func line(board [][]byte) bool {
	for i := 0; i < len(board); i++ {
		t := map[byte]int{}
		for j := 0; j < len(board[i]); j++ {
			t[board[i][j]]++
			n := t[board[i][j]]
			if n > 1 && board[i][j] != '.' {
				return false
			}
		}
	}
	return true
}

func row(board [][]byte) bool {
	for i := 0; i < len(board); i++ {
		t := map[byte]int{}
		for j := 0; j < len(board[i]); j++ {
			t[board[j][i]]++
			n := t[board[j][i]]
			if n > 1 && board[j][i] != '.' {
				return false
			}
		}
	}
	return true
}

func ninexnine(board [][]byte) bool {
	for i := 0; i < 3; i++ {
		for j := 0; j < 3; j++ {
			rStart := i * 3
			rEnd := i*3 + 3
			lStart := j * 3
			lEnd := j*3 + 3
			tMap := map[byte]int{}
			for t := rStart; t < rEnd; t++ {
				for n := lStart; n < lEnd; n++ {
					tMap[board[t][n]]++
					g := tMap[board[t][n]]
					if g > 1 && board[t][n] != '.' {
						return false
					}
				}
			}
		}
	}

	return true
}
