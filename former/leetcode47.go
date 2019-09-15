package main

import (
	"fmt"
	"sort"
)

func permuteUnique(nums []int) [][]int {
	res := [][]int{}
	if len(nums) == 0 {
		return res
	}
	sort.Ints(nums)
	res = append(res, nums[:1])
	for _, num := range nums[1:] {
		res = add2(res, num)
	}
	return res
}

func add2(r [][]int, num int) [][]int {
	res := [][]int{}
	for i := 0; i <= len(r[0]); i++ {
		for _, nn := range r {
			if i != 0 && num == nn[i-1] {
				continue
			} else {
				tmp := append([]int{}, nn...)
				tp := []int{}
				tp = append(tmp[:i], num)
				tp = append(tp, nn[i:]...)
				res = append(res, tp)
			}
		}
	}
	return res
}

func main() {
	nums := []int{1, 1, 2, 2}
	fmt.Println(permuteUnique(nums))
}
