package main

import "fmt"

func insert(intervals [][]int, newInterval []int) [][]int {
	newi := make([][]int, len(intervals)+1)
	for i := 0; i < len(intervals)-1; i++ {
		if intervals[i][0] <= newInterval[0] && newInterval[0] < intervals[i+1][0] {
			newi = append([][]int{}, newInterval)
			newi = append(newi, intervals[i+1:]...)
			newi = append(intervals[:i+1], newi...)
			//fmt.Println(intervals)
			break
		}
	}
	if len(intervals) == 0 {
		newi = [][]int{newInterval}
	} else {
		if len(intervals) == 1 {
			if intervals[0][0] < newInterval[0] {
				newi = append(intervals, newInterval)
			} else {
				newi = append([][]int{newInterval}, intervals...)
			}
		} else {
			if newInterval[0] < intervals[0][0] {
				newi = append([][]int{newInterval}, intervals...)
			} else {
				newi = append(intervals, newInterval)
			}
		}
	}

	//fmt.Println(newi)
	for i := 0; i < len(newi)-1; i++ {
		if newi[i+1][0] <= newi[i][1] {
			newi[i+1][0] = newi[i][0]
			if newi[i+1][1] < newi[i][1] {
				newi[i+1][1] = newi[i][1]
			}
			newi = append(newi[:i], newi[i+1:]...)
			i--
		}
	}

	return newi
}

func main() {
	intervals := [][]int{
		//[1,2],[3,5],[6,7],[8,10],[12,16]],
		/*[]int{1, 2},
		[]int{3, 5},
		[]int{6, 7},
		[]int{8, 10},
		[]int{12, 16},*/
		//[]int{1, 5},
		[]int{2, 6},
		[]int{7, 9},
	}
	newInterval := []int{15, 18}
	fmt.Println(insert(intervals, newInterval))
}
