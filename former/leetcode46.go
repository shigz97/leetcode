package main

import "fmt"

func permute(nums []int) [][]int {
	res := [][]int{}
	if len(nums) == 0 {
		return res
	}
	res = append(res, nums[:1])
	for _, num := range nums[1:] {
		res = add(res, num)
	}
	return res
}

func add(r [][]int, num int) [][]int {
	res := [][]int{}
	for i := 0; i <= len(r[0]); i++ {
		for _, nn := range r {
			tmp := append([]int{}, nn...)
			tp := []int{}
			tp = append(tmp[:i], num)
			tp = append(tp, nn[i:]...)
			res = append(res, tp)
		}
	}
	return res
}

func main() {
	nums := []int{1, 2, 3, 4}
	fmt.Println(len(permute(nums)))
}
