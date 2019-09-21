package main

import "fmt"

func subsets(nums []int) [][]int {
	res := [][]int{}
	res = append(res, []int{})
	for i := range nums {
		addslice(i, []int{}, nums, &res)
	}
	return res
}

func addslice(i int, tmp, nums []int, res *[][]int) {
	tp := append([]int{}, tmp...)
	tp = append(tp, nums[i])
	*res = append(*res, tp)
	for j := i + 1; j < len(nums); j++ {
		addslice(j, tp, nums, res)
	}
}

func main() {
	nums := []int{1, 2, 3, 4, 5}
	fmt.Println(subsets(nums))
}
