package main

import (
	"fmt"
	"sort"
)

func majorityElement(nums []int) int {
	sort.Ints(nums)
	return nums[len(nums)/2]
}

func main() {
	nums := []int{3, 2, 3}
	fmt.Println(majorityElement(nums))
}
