package main

import "fmt"

func containsDuplicate(nums []int) bool {
	dic := make(map[int]struct{}, len(nums))
	var mem = struct{}{}
	for _, n := range nums {
		if _, ok := dic[n]; ok {
			return false
		}
		dic[n] = mem
	}
	return true
}

func main() {
	nums := []int{1, 2, 3, 1}
	fmt.Println(containsDuplicate(nums))
}
