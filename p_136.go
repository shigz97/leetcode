package main

import "fmt"

func singleNumber(nums []int) int {
	res := 0
	for _, n := range nums {
		res = res ^ n
	}
	return res
}

func main() {
	nums := []int{4, 1, 2, 1, 2}
	fmt.Println(singleNumber(nums))
}
