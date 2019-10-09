//用堆
package main

import "fmt"

func findKthLargest(nums []int, k int) int {
	buildHeap(nums)
	if k == 1 {
		return nums[0]
	} else {
		fmt.Println(nums[0])
		return findKthLargest(nums[1:], k-1)
	}
}

func buildHeap(nums []int) {
	//fmt.Println(nums)
	for i := (len(nums) - 2) / 2; i >= 0; i-- {
		//tmp := nums[i]
		j := i*2 + 1
		if j+1 < len(nums) && nums[j+1] > nums[j] {
			j = j + 1
		}
		if nums[i] < nums[j] {
			nums[i], nums[j] = nums[j], nums[i]
		}
	}
}

func main() {
	nums := []int{1}
	fmt.Println(findKthLargest(nums, 1))
}
