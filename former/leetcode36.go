package main

import "fmt"

func isValidSudoku(board [][]byte) bool {
	res := true
	for _, str := range board {
		pos := make([]int, 10)
		for _, b := range str {
			if b != '.' {
				if pos[b-48] == 1 {
					res = false
					return res
				} else {
					pos[b-48] = 1
				}
			}
		}
	}
	for i := 0; i < 9; i++ {
		pos := make([]int, 10)
		for j := 0; j < 9; j++ {
			if board[j][i] != '.' {
				if pos[board[j][i]-48] == 1 {
					res = false
					return res
				} else {
					pos[board[j][i]-48] = 1
				}
			}
		}
	}
	for i := 1; i < 8; i += 3 {
		for j := 1; j < 8; j += 3 {
			pos := make([]int, 10)
			for _, k := range []int{i - 1, i, i + 1} {
				for _, t := range []int{j - 1, j, j + 1} {
					if board[k][t] != '.' {
						if pos[board[k][t]-48] == 1 {
							res := false
							return res
						} else {
							pos[board[k][t]-48] = 1
						}
					}
				}
			}
		}
	}
	return res
}

func main() {
	board := [][]byte{
		[]byte("53..7...."),
		[]byte("6..195..."),
		[]byte(".98....6."),
		[]byte("8...6...3"),
		[]byte("4..8.3..1"),
		[]byte("7...2...6"),
		[]byte(".6....28."),
		[]byte("...419..5"),
		[]byte("....8..79"),
	}
	fmt.Println(isValidSudoku(board))
}
