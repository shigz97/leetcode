package main

import (
	"fmt"
	"math"
)

func rob(nums []int) int {
	pre, cur := 0, 0
	for _, n := range nums {
		pre, cur = cur, int(math.Max((float64(pre+n)), float64(cur)))
	}
	return cur
}

func main() {
	nums := []int{1, 2, 3, 1}
	fmt.Println(rob(nums))
}
