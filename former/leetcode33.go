package main

import "fmt"

func search(nums []int, target int) int {
	return searchh(nums, target, 0, len(nums)-1)
}

func searchh(nums []int, target int, l int, r int) int {
	if l == r && nums[l] != target {
		return -1
	}
	c := (l + r) / 2
	//fmt.Println(c, nums[c])
	if target == nums[c] {
		return c
	}
	if nums[l] < nums[c] {
		if target >= nums[l] && target <= nums[c] {
			return searchh(nums, target, l, c-1)
		} else {
			return searchh(nums, target, c+1, r)
		}
	} else {
		if target >= nums[c] && target <= nums[r] {
			return searchh(nums, target, c+1, r)
		} else {
			return searchh(nums, target, l, c-1)
		}
	}
}

func main() {
	nums := []int{4, 5, 6, 7, 0, 1, 2}
	fmt.Println(search(nums, 3))
}
