package main

import (
	"fmt"
)

func firstMissingPositive(nums []int) int {
	if len(nums) == 0 {
		return 1
	}
	nums = append(nums, 0)
	size := len(nums)
	for i := range nums {
		if nums[i] <= 0 || nums[i] >= size {
			nums[i] = 0
		}
	}
	for i := range nums {
		var loc int
		if nums[i] < 0 {
			loc = -nums[i]
		} else {
			loc = nums[i]
		}
		if loc < size {
			switch {
			case nums[loc] == 0:
				nums[loc] = -size - 1
			case nums[loc] > 0:
				nums[loc] = -nums[loc]
			}
		}
	}
	fmt.Println(nums)
	res := 1
	for res < size && nums[res] < 0 {
		res++
	}
	return res
}

func main() {
	nums := []int{0}
	fmt.Println(firstMissingPositive(nums))
}
