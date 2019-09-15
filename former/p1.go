package main

import "fmt"

func Twosum(nums []int, target int) []int {
	index := make(map[int]int, len(nums))
	for i, b := range nums {
		j, ok := index[target-b]
		if ok {
			return []int{j, i}
		}
		index[b] = i
	}
	return nil
}

func p1() {
	var num = []int{2, 7, 11, 13}
	var target int = 15
	fmt.Println(Twosum(num, target))
}
