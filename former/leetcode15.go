package main

import (
	"fmt"
	"sort"
)

func threeSum(nums []int) [][]int {
	/*res := [][]int{}
	if len(nums) == 0 {
		return res
	}
	sort.Ints(nums)
	max := nums[len(nums)-1]
	ii := 0
	if nums[0] > 0 {
		return res
	}
	for nums[ii]+max*2 < 0 && ii < len(nums)-1 {
		ii++
	}
	for i := ii; i < len(nums); i++ {
		if i > 0 && nums[i] == nums[i-1] {
			continue
		}
		dic := make(map[int][]int, len(nums[i:]))
		for j := i + 1; j < len(nums); j++ {
			if nums[i]+nums[j]+max < 0 {
				continue
			}
			if k, ok := dic[nums[j]]; ok {
				tmp := append(k, nums[j])
				res = append(res, tmp)
				for j+1 < len(nums) && nums[j+1] == nums[j] {
					j++
				}
			} else {
				if nums[j] == nums[j-1] && j-1 != i {
					//j = j + 1
					continue
				} else {
					t := -(nums[i] + nums[j])
					if t < nums[j] {
						continue
					} else {
						dic[t] = []int{nums[i], nums[j]}
					}
				}
			}
			fmt.Println(j, dic)
		}
	}
	return res*/
	res := [][]int{}
	if len(nums) == 0 {
		return res
	}
	sort.Ints(nums)
	for i := range nums {
		if i > 0 && nums[i] == nums[i-1] {
			continue
		}

		l, r := i+1, len(nums)-1

		for l < r {
			sum := nums[i] + nums[l] + nums[r]
			if sum > 0 {
				r--
			} else if sum < 0 {
				l++
			} else {
				tmp := []int{nums[i], nums[l], nums[r]}
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
	return res
}

func main() {
	nums := []int{-4, -2, -2, -2, 0, 1, 2, 2, 2, 3, 3, 4, 4, 6, 6}
	fmt.Println(threeSum(nums))
}
