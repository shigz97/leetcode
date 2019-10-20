package main

import "fmt"

func checkStraightLine(coordinates [][]int) bool {
	k := (float64)(coordinates[1][1]-coordinates[0][1]) / (float64)(coordinates[1][0]-coordinates[0][0])
	fmt.Println(k)
	for _, c := range coordinates[2:] {
		t := (float64)(c[1]-coordinates[0][1]) / (float64)(c[0]-coordinates[0][0])
		if t != k {
			return false
		}
	}
	return true
}

func main() {
	nums := [][]int{
		[]int{1, 1},
		[]int{2, 2},
		[]int{3, 4},
		[]int{4, 5},
		[]int{5, 6},
		[]int{7, 7},
	}
	fmt.Println(checkStraightLine(nums))
}
