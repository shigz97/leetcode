package main

import (
	"fmt"
	"sort"
)

func threeSumClosest(nums []int, target int) int {
	sort.Ints(nums)
	res := nums[0] + nums[1] + nums[2]
	for i, m := range nums {
		l, r := i+1, len(nums)-1
		for l < r {
			sum := nums[l] + nums[r]
			if abs(sum+m-target) < abs(res-target) {
				res = sum + m
			}
			switch {
			case sum+m < target:
				l++
			case sum+m > target:
				r--
			default:
				return target
			}
		}
	}
	return res
}

func abs(a int) int {
	if a < 0 {
		return -a
	}
	return a
}

func main() {
	nums := []int{-1, 2, 1, 4}
	target := 1
	fmt.Println(threeSumClosest(nums, target))
}
