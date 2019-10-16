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
	fmt.Println(start)
	if _, ok := dic[start]; ok {
		return
	} else {
		//dic[start] = mem
	}
	switch dir {
	case 0:
		for start.j < len(matrix[0]) {
			if _, ok := dic[start]; ok {
				break
			}
			fmt.Println(matrix[start.i][start.j])
			*res = append(*res, matrix[start.i][start.j])
			start.j = start.j + 1
		}
		start.i, start.j = start.i+1, start.j-1
		findSpiral(matrix, start, 1, res)
		return
	case 1:
		for start.i < len(matrix) {
			if _, ok := dic[start]; ok {
				break
			}
			*res = append(*res, matrix[start.i][start.j])
			start.i = start.i + 1
		}
		start.i, start.j = start.i-1, start.j-1
		findSpiral(matrix, start, 2, res)
		return
	case 2:
		return
	case 3:
		return
	}
	return
}

func main() {
	matrix := [][]int{
		[]int{1, 2, 3},
		[]int{4, 5, 6},
		[]int{7, 8, 9},
	}
	fmt.Println(spiralOrder(matrix))
}
