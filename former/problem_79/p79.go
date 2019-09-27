package main

import "fmt"

func exist(board [][]byte, word string) bool {
	res := false
	if len(board) == 0 {
		return res
	}
	m, n := len(board), len(board[0])
	isVisited := make([][]int, m)
	for i := range isVisited {
		isVisited[i] = make([]int, n)
	}
	w := []byte(word)
	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			if board[i][j] == w[0] {
				res = res || judge(board, isVisited, m, n, i, j, w)
			}
		}
	}
	return res
}

func judge(board [][]byte, isVisited [][]int, m, n, i, j int, w []byte) bool {
	if len(w) == 0 {
		return true
	}
	if i < 0 || i >= m {
		return false
	}
	if j < 0 || j >= n {
		return false
	}
	if isVisited[i][j] == 1 {
		return false
	}
	isVisited[i][j] = 1
	fmt.Println(i, j, isVisited)
	return w[0] == board[i][j] && (judge(board, isVisited, m, n, i-1, j, w[1:]) || judge(board, isVisited, m, n, i, j+1, w[1:]) ||
		judge(board, isVisited, m, n, i+1, j, w[1:]) || judge(board, isVisited, m, n, i, j-1, w[1:]))
}

func main() {
	/*board := [][]byte{
		[]byte{'A', 'B', 'C', 'E'},
		[]byte{'S', 'F', 'C', 'S'},
		[]byte{'A', 'D', 'E', 'E'},
	}*/
	board := [][]byte{
		[]byte{'a', 'b'},
		[]byte{'c', 'd'},
	}
	fmt.Println(exist(board, "acbd"))
}
