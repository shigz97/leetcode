package main

import "fmt"

//暴力解法
/*func maxArea(height []int) int {
	max := 0
	for i, k := range height {
		for j, p := range height[i+1:] {
			if getmin(k, p)*(j+1) > max {
				max = getmin(k, p) * (j + 1)
			}
		}
	}
	return max
}
func getmin(a, b int) int {
	if a > b {
		return b
	} else {
		return a
	}
}*/

func main() {
	var nums = []int{1, 8, 6, 2, 5, 4, 8, 3, 7}
	fmt.Println(maxArea(nums))
}
