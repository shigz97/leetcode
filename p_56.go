package main

import (
	"fmt"
	"sort"
)

type myint [][]int

func (m myint) Len() int {
	return len(m)
}

func (m myint) Less(i, j int) bool {
	if m[i][0] < m[j][0] {
		return true
	}
	return false
}

func (m myint) Swap(i, j int) {
	m[i], m[j] = m[j], m[i]
}

func merge(intervals [][]int) [][]int {
	var inter myint = intervals
	sort.Sort(inter)
	for i := 0; i < len(inter)-1; i++ {
		if inter[i+1][0] <= inter[i][1] {
			inter[i+1][0] = inter[i][0]
			if inter[i+1][1] < inter[i][1] {
				inter[i+1][1] = inter[i][1]
			}
			inter = append(inter[:i], inter[i+1:]...)
			i--
		}
	}
	return inter
}

func main() {
	intervals := [][]int{
		[]int{1, 3},
		[]int{2, 6},
		[]int{8, 10},
		[]int{15, 18},
	}
	fmt.Println(merge(intervals))
}
