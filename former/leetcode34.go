package main

import "fmt"

func searchRange(nums []int, target int) []int {
	if len(nums) == 0 {
		return []int{-1, -1}
	}
	if len(nums) == 1 {
		if nums[0] == target {
			return []int{0, 0}
		} else {
			return []int{-1, -1}
		}
	}
	l, mid, r := 0, 0, len(nums)-1
	found := false
	for l <= r {
		mid = (l + r) / 2
		fmt.Println(l, r, mid, nums[mid])
		if nums[mid] == target {
			found = true
			break
		} else {
			if nums[mid] > target {
				r = mid - 1
			} else {
				l = mid + 1
			}
		}
	}
	if found == false {
		return []int{-1, -1}
	} else {
		s, e := mid-1, mid+1
		for s >= 0 && nums[s] == nums[mid] {
			s--
		}
		for e < len(nums) && nums[e] == nums[mid] {
			e++
		}
		return []int{s + 1, e - 1}
	}
	return []int{-1, -1}
}

func main() {
	nums := []int{1, 4}
	fmt.Println(searchRange(nums, 4))
}
