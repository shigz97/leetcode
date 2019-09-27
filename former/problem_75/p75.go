package main

import "fmt"

func sortColors(nums []int) {
	p0, p2 := 0, len(nums)-1
	for nums[p0] != 2 && p0 <= len(nums)-1 {
		p0++
	}
	for nums[p2] != 0 && p2 > 0 {
		p2--
	}
	cur := 0
	for cur <= p2 {
		switch {
		case nums[cur] == 0:
			nums[cur], nums[p0] = nums[p0], nums[cur]
			cur++
		case nums[cur] == 1:
			cur++
		case nums[cur] == 2:
			nums[cur], nums[p2] = nums[p2], nums[cur]
			cur, p2 = cur+1, p2-1
		}
	}
}

func main() {
	nums := []int{2, 0, 1}
	sortColors(nums)
	fmt.Println(nums)
}
