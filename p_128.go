package main

import "fmt"

func longestConsecutive(nums []int) int {
	count := make(map[int]int, len(nums))
	max := 0
	for _, n := range nums {
		if count[n] != 0 {
			continue
		}
		count[n] = 1 + count[n-1] + count[n+1]
		if count[n] > max {
			max = count[n]
		}
		count[n-count[n-1]] = count[n]
		count[n+count[n+1]] = count[n]
	}
	return max
}

func main() {
	nums := []int{5, 100, 4, 200, 1, 3, 2}
	fmt.Println(longestConsecutive(nums))
}
