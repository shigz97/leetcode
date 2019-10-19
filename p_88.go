package main

import (
	"fmt"
	"sort"
)

func merge(nums1 []int, m int, nums2 []int, n int) {
	for i, j := m, 0; j < n; i, j = i+1, j+1 {
		nums1[i] = nums2[j]
	}
	sort.Ints(nums1[:m+n])
	fmt.Println(nums1)
}

func main() {
	nums1 := []int{1, 2, 3, 0, 0, 0, 0}
	nums2 := []int{2, 5, 6}
	merge(nums1, 3, nums2, 3)
}
