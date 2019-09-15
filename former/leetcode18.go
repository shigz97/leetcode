package main

import (
	"fmt"
	"sort"
)

func fourSum(nums []int, target int) [][]int {
	if len(nums) == 0 {
		return [][]int{}
	}
	sort.Ints(nums)
	//max := nums[len(nums)-1]
	res := [][]int{}
	for i := 1; i < len(nums); i++ {
		if i > 1 && nums[i] == nums[i-1] && nums[i-1] == nums[i-2] {
			continue
		}
		edge := 0
		if nums[i] == nums[i-1] {
			edge = i - 1
		}
		for j := i - 1; j >= edge; j-- {
			if j+1 != i && nums[j] == nums[j+1] {
				continue
			}
			l, r := i+1, len(nums)-1
			for l < r {
				sum := nums[j] + nums[i] + nums[l] + nums[r]
				switch {
				case sum > target:
					r--
				case sum < target:
					l++
				default:
					tmp := []int{nums[j], nums[i], nums[l], nums[r]}
					res = append(res, tmp)
					for l < r {
						l++
						if nums[l] != nums[l-1] {
							break
						}
					}
					for l < r {
						r--
						if nums[r] != nums[r+1] {
							break
						}
					}
				}
			}
		}
	}
	return res
}

func main() {
	nums := []int{-3, -2, -1, 0, 0, 1, 2, 3}
	fmt.Println(fourSum(nums, 0))
}
