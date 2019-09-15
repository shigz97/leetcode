package main

import "fmt"

func main() {
	n1 := []int{1, 2}
	n2 := []int{3, 4}
	fmt.Println(findMedianSortedArrays(n1, n2))
}

func findMedianSortedArrays(nums1 []int, nums2 []int) float64 {
	var vec []int
	var i, j int
	i, j = 0, 0
	for i < len(nums1) && j < len(nums2) {
		if nums1[i] < nums2[j] {
			vec = append(vec, nums1[i])
			i++
		} else {
			vec = append(vec, nums2[j])
			j++
		}
	}
	vec = append(vec, nums1[i:]...)
	vec = append(vec, nums2[j:]...)
	//fmt.Printf("vec=%v", vec)
	loc := len(vec)
	if loc%2 == 1 {
		return float64(vec[loc/2])
	} else {
		var res float64
		res = (float64(vec[loc/2-1]) + float64(vec[loc/2])) / 2
		return res
	}
	return 0.0
}
