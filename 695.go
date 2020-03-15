package main

import "fmt"

func main() {
	grid := [][]int{
		[]int{0, 0, 1, 0, 0, 0, 0, 1, 0, 0, 0, 0, 0},
		[]int{0, 0, 0, 0, 0, 0, 0, 1, 1, 1, 0, 0, 0},
		[]int{0, 1, 1, 0, 1, 0, 0, 0, 0, 0, 0, 0, 0},
		[]int{0, 1, 0, 0, 1, 1, 0, 0, 1, 0, 1, 0, 0},
		[]int{0, 1, 0, 0, 1, 1, 0, 0, 1, 1, 1, 0, 0},
		[]int{0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 1, 0, 0},
		[]int{0, 0, 0, 0, 0, 0, 0, 1, 1, 1, 0, 0, 0},
		[]int{0, 0, 0, 0, 0, 0, 0, 1, 1, 0, 0, 0, 0},
	}
	fmt.Println(maxAreaOfIsland(grid))
}

func maxAreaOfIsland(grid [][]int) int {
	m, n := len(grid), len(grid[0])
	max := 0
	cur := 0
	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			if grid[i][j] == 1 {
				cur = 0
				dfs(grid, i, j, &max, &cur, m, n)
			}
		}
	}
	return max
}

func dfs(grid [][]int, i, j int, max *int, cur *int, m, n int) {
	if i >= 0 && i < m && j >= 0 && j < n && grid[i][j] == 1 {
		*cur++
		if *cur >= *max {
			//fmt.Println(i, j, *cur)
			*max = *cur
		}
		grid[i][j] = 0
		dfs(grid, i-1, j, max, cur, m, n)
		dfs(grid, i, j+1, max, cur, m, n)
		dfs(grid, i+1, j, max, cur, m, n)
		dfs(grid, i, j-1, max, cur, m, n)
	}
}
