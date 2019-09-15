package main

import (
	"fmt"
	"sort"
)

func combinationSum(candidates []int, target int) [][]int {
	sort.Ints(candidates)
	res := [][]int{}
	tmp := []int{}
	maxsize := target/candidates[0] + 1
	stack, top := make([]int, maxsize), -1
	i := 0
	for top > -1 || (i < len(candidates) && candidates[i] <= target) {
		if i < len(candidates) && candidates[i] <= target {
			if candidates[i] < target {
				tmp = append(tmp, candidates[i])
				target = target - candidates[i]
				top++
				stack[top] = i
			} else {
				tmp = append(tmp, candidates[i])
				tp := []int{}
				tp = append(tp, tmp...)
				res = append(res, tp)
				if top != -1 {
					tp := stack[top]
					top--
					target = target + candidates[tp]
					i = tp + 1
					tmp = tmp[:len(tmp)-2]
				} else {
					break
				}
			}
		} else {
			tp := stack[top]
			top--
			target = target + candidates[tp]
			i = tp + 1
			tmp = tmp[:len(tmp)-1]
		}
	}
	return res
}

func main() {
	candidates := []int{2, 3, 5}
	target := 8
	fmt.Println(combinationSum(candidates, target))
}
