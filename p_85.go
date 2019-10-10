package main

import "fmt"

func maximalRectangle(matrix [][]byte) int {
	x, y := len(matrix), len(matrix[0])
	width := make([][]int, x)
	for i, m := range matrix {
		width[i] = make([]int, y)
		for j, n := range m {
			if n == '1' {
				if j == 0 {
					width[i][j] = 1
				} else {
					width[i][j] = width[i][j-1] + 1
				}
			} else {
				width[i][j] = 0
			}
		}
	}
	//fmt.Println(width)
	max := 0
	for i := 0; i < x; i++ {
		for j := 0; j < y; j++ {
			if width[i][j] != 0 {
				wid := width[i][j]
				var k int
				for k = i; k >= 0 && width[k][j] != 0; k-- {
					if width[k][j] < wid {
						wid = width[k][j]
					}
					area := wid * (i - k + 1)
					if area > max {
						max = area
					}
				}
				//fmt.Println(i, j, k, wid)
			}
		}
	}
	return max
}

/*
'1','0','1','0','0'
'1','0','1','1','1'
'1','1','1','1','1'
'1','0','0','1','0'
*/
func main() {
	matrix := [][]byte{
		[]byte{'0', '1', '1', '0', '1'},
		[]byte{'1', '1', '0', '1', '0'},
		[]byte{'0', '1', '1', '1', '0'},
		[]byte{'1', '1', '1', '1', '0'},
		[]byte{'1', '1', '1', '1', '1'},
		[]byte{'0', '0', '0', '0', '0'},
	}
	fmt.Println(maximalRectangle(matrix))
}
