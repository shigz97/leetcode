package main

import "fmt"

func maxProfit(prices []int) int {
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
	p := []int{7, 1, 5, 3, 6, 4}
	fmt.Println(maxProfit(p))
}
