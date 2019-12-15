package main

import "fmt"

func setZeroes(matrix [][]int) {
	if len(matrix) == 0 {
		return
	}
	m, n := len(matrix), len(matrix[0])
	zeros := make([]int, n)
	for i := 0; i < m; i++ {
		index := getIndex(matrix[i])
		if len(index) != 0 {
			for j := 0; j < n; j++ {
				matrix[i][j] = 0
			}
		}
		for _, d := range index {
			if zeros[d] == 0 {
				zeros[d] = 1
			}
		}
	}
	//fmt.Println(zeros)
	for j, z := range zeros {
		if z != 0 {
			for i := 0; i < m; i++ {
				matrix[i][j] = 0
			}
		}
	}
}

func getIndex(in []int) []int {
	res := []int{}
	for i := range in {
		if in[i] == 0 {
			res = append(res, i)
		}
	}
	return res
}

func main() {
	matrix := [][]int{
		[]int{0, 1, 2, 0},
		[]int{3, 4, 5, 2},
		[]int{1, 3, 1, 5},
	}
	setZeroes(matrix)
	fmt.Println(matrix)
}
