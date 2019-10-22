package main

import (
	"fmt"
	"sort"
)

func maxProfit(prices []int) int {
	if len(prices) == 0 {
		return 0
	}
	i, j := 0, 1
	res := []int{}
	for j < len(prices) {
		if prices[j-1] > prices[j] {
			res = append(res, maxSingleProfit(prices[i:j]))
			i = j
		}
		j++
	}
	res = append(res, maxSingleProfit(prices[i:j]))
	//sort.Reverse(sort.Ints(res))
	sort.Sort(sort.Reverse(sort.IntSlice(res)))
	fmt.Println(res)
	if len(res) >= 2 {
		return res[0] + res[1]
	} else {
		return res[0]
	}
}

func maxSingleProfit(prices []int) int {
	res, min := 0, -1
	for _, p := range prices {
		if min == -1 || min > p {
			min = p
		} else {
			if p-min > res {
				res = p - min
			}
		}
	}
	return res
}

func main() {
	p := []int{1, 2, 4, 2, 5, 7, 2, 4, 9, 0}
	fmt.Println(maxProfit(p))
}
