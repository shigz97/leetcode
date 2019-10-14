package main

import "fmt"

func maxSlidingWindow(nums []int, k int) []int {
	if len(nums) == 0 {
		return []int{}
	}
	res := make([]int, len(nums)-k+1)
	maxloc := -1
	for i := 0; i < len(res); i++ {
		if maxloc == -1 || maxloc < i {
			maxloc = i + getMaxLoc(nums[i:i+k])
		} else {
			if nums[maxloc] < nums[i+k-1] {
				maxloc = i + k - 1
			}
		}
		res[i] = nums[maxloc]
	}
	return res
}

func getMaxLoc(nums []int) int {
	res := 0
	for i, n := range nums[1:] {
		if n > nums[res] {
			res = i + 1
		}
	}
	return res
}

func main() {
	nums := []int{1, 3, -1, -3, 5, 3, 6, 7}
	fmt.Println(maxSlidingWindow(nums, 2))
}
