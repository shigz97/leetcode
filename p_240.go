package main

import "fmt"

func searchMatrix(matrix [][]int, target int) bool {
	if len(matrix) == 0 || len(matrix[0]) == 0 {
		return false
	}
	return isMatrixLoc(matrix, 0, 0, target)
}

func isMatrixLoc(matrix [][]int, i, j int, target int) bool {
	fmt.Println(i, j, matrix[i][j])
	if matrix[i][j] == target {
		return true
	}
	var res bool
	if matrix[i][j] > target || j == len(matrix[0])-1 {
		res = false
	} else {
		res = res || isMatrixLoc(matrix, i, j+1, target)
	}
	if res == true {
		return true
	}
	if i == len(matrix)-1 {
		return false
	}
	res = res || isMatrixLoc(matrix, i+1, j, target)
	return res
}

func main() {
	matrix := [][]int{
		[]int{1, 4, 7, 11, 15},
		[]int{2, 5, 8, 12, 19},
		[]int{3, 6, 9, 16, 22},
		[]int{10, 13, 14, 17, 24},
		[]int{18, 21, 23, 26, 30},
	}
	//fmt.Println(matrix)
	fmt.Println(searchMatrix(matrix, 20))
}
