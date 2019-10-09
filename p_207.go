package main

import "fmt"

func canFinish(numCourses int, prerequisites [][]int) bool {
	//isVisited := make([]int, numCourses)
	if len(prerequisites) == 0 {
		return true
	}
	res := true
	for i := 0; i < numCourses; i++ {
		isVisited := make([]int, numCourses)
		res = res && dfs(prerequisites, isVisited, i)
		if res == false {
			return false
		}
	}

	return res
}

func check(isVisited []int) (int, bool) {
	for j, i := range isVisited {
		if i == 0 {
			return j, false
		}
	}
	return -1, true
}

//DFS
func dfs(pre [][]int, isVisited []int, node int) bool {
	isVisited[node] = 1
	//tmp := append([]int{}, isVisited...)
	//fmt.Println(isVisited)
	res := true
	for _, p := range pre {
		if p[0] == node {
			if isVisited[p[1]] == 1 {
				return false
			} else {
				res = res && dfs(pre, isVisited, p[1])
				isVisited[p[1]] = 0
			}
		}
	}
	return res
}

func main() {
	pre := [][]int{
		[]int{2, 0},
		[]int{1, 0},
		[]int{3, 1},
		[]int{3, 2},
		[]int{1, 3},
	}
	fmt.Println(canFinish(4, pre))
}
