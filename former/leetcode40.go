package main

import (
	"fmt"
	"sort"
)

func combinationSum2(candidates []int, target int) [][]int {
	res := [][]int{}
	if len(candidates) == 0 {
		return res
	}
	tmp := []int{}
	sort.Ints(candidates)
	maxlen := target/candidates[0] + 1
	stack, top := make([]int, maxlen), -1
	i := 0
	for top > -1 || (i < len(candidates) && candidates[i] <= target) {
		if i < len(candidates) && candidates[i] <= target {
			if candidates[i] < target {
				tmp = append(tmp, candidates[i])
				target = target - candidates[i]
				top++
				stack[top] = i
				i++
			} else {
				tmp = append(tmp, candidates[i])
				tp := []int{}
				tp = append(tp, tmp...)
				res = append(res, tp)
				if top != -1 {
					tp := stack[top]
					top--
					target = target + candidates[tp]
					for i = tp + 1; i < len(candidates) && candidates[i] == candidates[tp]; i++ {
					}
					tmp = tmp[:len(tmp)-2]
				} else {
					break
				}
			}
		} else {
			tp := stack[top]
			top--
			target = target + candidates[tp]
			for i = tp + 1; i < len(candidates) && candidates[i] == candidates[tp]; i++ {
			}
			tmp = tmp[:len(tmp)-1]
		}
	}
	return res
}

func main() {
	candidates := []int{1, 1}
	target := 2
	fmt.Println(combinationSum2(candidates, target))
}
