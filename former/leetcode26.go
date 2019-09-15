package main

import "fmt"

func removeDuplicates(nums []int) int {
	if len(nums) == 0 {
		return 0
	}
	former := nums[0]
	count := 0
	for i := 1; i < len(nums); i++ {
		if nums[i] == former {
			count++
			nums = append(nums[:i], nums[i+1:]...)
			i--
		} else {
			former = nums[i]
		}
	}
	fmt.Println(nums)
	return count
}

func main() {
	nums := []int{1, 1, 1, 2, 2, 2}
	fmt.Println(removeDuplicates(nums))
}
