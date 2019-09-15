package main

import "fmt"

func removeElement(nums []int, val int) int {
	if len(nums) == 0 {
		return 0
	}
	i, j := 0, 0
	for j < len(nums) {
		if nums[j] == val {
			for j < len(nums) && nums[j] == val {
				j++
			}
			if j < len(nums) {
				nums[i] = nums[j]
				i, j = i+1, j+1
			}
		} else {
			nums[i] = nums[j]
			i, j = i+1, j+1
		}
		fmt.Println(i, j, nums)
	}
	return j - i
}

func main() {
	nums := []int{0, 1, 2, 2, 3, 0, 4, 2}
	fmt.Println(removeElement(nums, 2))
}
