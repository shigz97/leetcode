package main

import "fmt"

type snode struct {
	i, j int
}

var dic = make(map[snode]struct{})
var mem = struct{}{}

func spiralOrder(matrix [][]int) []int {
	res := []int{}
	start := snode{0, 0}
	findSpiral(matrix, start, 0, &res)
	return res
}

func findSpiral(matrix [][]int, start snode, dir int, res *[]int) {
	//i,j := start.i,start.j
	//fmt.Println(start)
	if _, ok := dic[start]; ok {
		return
	} else {
		dic[start] = mem
	}
	//fmt.Println(start)
	switch dir {
	case 0:
		for start.j < len(matrix[0]) {
			*res = append(*res, matrix[start.i][start.j])
			start.j = start.j + 1
			if _, ok := dic[start]; ok {
				break
			}
		}
		start.i, start.j = start.i+1, start.j-1
		findSpiral(matrix, start, 1, res)
		return
	case 1:
		for start.i < len(matrix) {
			*res = append(*res, matrix[start.i][start.j])
			start.i = start.i + 1
			if _, ok := dic[start]; ok {
				break
			}
		}
		start.i, start.j = start.i-1, start.j-1
		findSpiral(matrix, start, 2, res)
		return
	case 2:
		for start.j >= 0 {
			*res = append(*res, matrix[start.i][start.j])
			start.j = start.j - 1
			if _, ok := dic[start]; ok {
				break
			}
		}
		start.i, start.j = start.i-1, start.j+1
		findSpiral(matrix, start, 3, res)
		return
	case 3:
		for start.i >= 0 {
			*res = append(*res, matrix[start.i][start.j])
			start.i = start.i - 1
			if _, ok := dic[start]; ok {
				break
			}
		}
		start.i, start.j = start.i+1, start.j+1
		findSpiral(matrix, start, 0, res)
		return
	}
	return
}

func main() {
	matrix := [][]int{
		[]int{1, 2, 3, 4},
		[]int{5, 6, 7, 8},
		[]int{9, 10, 11, 12},
		//[]int{13, 14, 15, 16},
	}
	fmt.Println(spiralOrder(matrix))
}
