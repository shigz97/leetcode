package main

import "fmt"

func largestRectangleArea(heights []int) int {
	area := 0
	split_num(heights, &area)
	return area
}

//分治
func split_num(nums []int, res *int) {
	if len(nums) == 0 {
		return
	}
	if len(nums) == 1 {
		if nums[0] > *res {
			*res = nums[0]
		}
		return
	}
	loc := getMin(nums)
	area := len(nums) * nums[loc]
	if area > *res {
		*res = area
	}
	split_num(nums[0:loc], res)
	split_num(nums[loc+1:len(nums)], res)
}

func getMin(nums []int) int {
	res := -1
	for i, n := range nums {
		if res == -1 || nums[res] > n {
			res = i
		}
	}
	return res
}

func main() {
	nums := []int{2, 1, 5, 6, 2, 3}
	fmt.Println(largestRectangleArea(nums))
}
