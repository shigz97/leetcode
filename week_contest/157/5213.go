package main

func minCostToMoveChips(chips []int) int {
	//count := make(map[int]int)
	odd, even := 0, 0
	for _, k := range chips {
		if k%2 == 0 {
			even++
		} else {
			odd++
		}
	}
	if even > odd {
		return odd
	} else {
		return even
	}
}
