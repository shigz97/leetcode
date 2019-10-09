package main

import "fmt"

func findKthLargest(nums []int, k int) int {
	//fmt.Println(nums, k)
	tmp := nums[0]
	l, r := 0, len(nums)-1
	//fmt.Println(l, r, nums)
	for l < r {
		for nums[r] < tmp && l != r {
			r--
		}
		if l == r {
			break
		}
		nums[l], l = nums[r], l+1
		for nums[l] > tmp && l != r {
			l++
		}
		if l == r {
			break
		}
		nums[r], r = nums[l], r-1
		//fmt.Println(l, r, nums)
	}
	nums[l] = tmp
	//fmt.Println(l, r, nums)
	switch {
	case l == k-1:
		return nums[l]
	case l < k-1:
		return findKthLargest(nums[l+1:], k-1-l)
	case l > k-1:
		return findKthLargest(nums[:l], k)
	}
	return 0
}

func main() {
	nums := []int{3, 2, 3, 1, 2, 4, 5, 5, 6}
	fmt.Println(findKthLargest(nums, 4))
}
