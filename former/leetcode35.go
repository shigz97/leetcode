package main

import "fmt"

func searchInsert(nums []int, target int) int {
	if len(nums) == 0 {
		return 0
	}
	if target < nums[0] {
		return 0
	}
	if target > nums[len(nums)-1] {
		return len(nums)
	}
	return findInsert(nums, target, 0, len(nums)-1)
}

func findInsert(nums []int, target int, l int, r int) int {
	if r == l+1 {
		switch {
		case nums[l] == target:
			return l
		case nums[r] == target:
			return r
		default:
			return r
		}
	}
	mid := (l + r) / 2
	switch {
	case nums[mid] == target:
		return mid
	case nums[mid] < target:
		return findInsert(nums, target, mid, r)
	default:
		return findInsert(nums, target, l, mid)
	}
}

func main() {
	nums := []int{1, 3, 5, 6}
	fmt.Println(searchInsert(nums, 0))
}
