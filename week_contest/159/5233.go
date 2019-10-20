package main

import "fmt"

func jobScheduling(startTime []int, endTime []int, profit []int) int {
	isVisit := make([]int, len(startTime))
	max := 0
	for i := range startTime {
		if isVisit[i] == 0 {
			findmoney(startTime, endTime, profit, i, &max, 0, isVisit)
		}
	}
	return max
}

func findmoney(start []int, end []int, pro []int, s int, max *int, cur int, vis []int) {
	cur = cur + pro[s]
	//fmt.Println(s, cur)
	if cur > *max {
		*max = cur
	}
	for i := 0; i < len(start); i++ {

		if start[i] >= end[s] {
			vis[i] = 1
			findmoney(start, end, pro, i, max, cur, vis)
		}

	}
	//vis[s] = 1
}

func main() {
	start := []int{6, 15, 7, 11, 1, 3, 16, 2}
	end := []int{19, 18, 19, 16, 10, 8, 19, 8}
	profit := []int{2, 9, 1, 19, 5, 7, 3, 19}
	fmt.Println(jobScheduling(start, end, profit))
}
