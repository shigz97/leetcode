package main

import "fmt"

func canJump(nums []int) bool {
	memo := make([]int, len(nums))
	for i := range memo {
		memo[i] = -1
	}
	return dyJump(nums, memo, 0)
}

func dyJump(nums []int, memo []int, loc int) bool {
	res := false
	for step := nums[loc]; step > 0; step-- {
		new_loc := loc + step
		if new_loc >= len(nums)-1 {
			res = true
			break
		} else {
			if memo[new_loc] == -1 {
				tmp := dyJump(nums, memo, new_loc)
				if tmp {
					memo[new_loc] = 1
				} else {
					memo[new_loc] = 0
				}
			}
			if memo[new_loc] == 1 {
				res = true
				break
			}
		}
	}
	return res
}

func main() {
	nums := []int{3, 2, 1, 0, 4}
	fmt.Println(canJump(nums))
}
