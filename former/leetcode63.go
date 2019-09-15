package main

import "fmt"

func uniquePathsWithObstacles(obstacleGrid [][]int) int {
	memo := [][]int{}
	for i := 0; i < len(obstacleGrid); i++ {
		tmp := make([]int, len(obstacleGrid[0]))
		memo = append(memo, tmp)
	}
	return find(obstacleGrid, memo, 0, 0)
}

func find(grid [][]int, memo [][]int, m, n int) int {
	if m == len(grid)-1 && n == len(grid[0])-1 {
		return 0
	}
	if grid[m][n] == 1 {
		return 0
	}
	if m == len(grid)-1 {
		memo[m][n+1] = find(grid, memo, m, n+1)
		return memo[m][n+1]
	}
	if n == len(grid[0])-1 {
		memo[m+1][n] = find(grid, memo, m+1, n)
		return memo[m][n]
	}
	if memo[m+1][n] == 0 {
		memo[m+1][n] = find(grid, memo, m+1, n)
	}
	if memo[m][n+1] == 0 {
		memo[m][n+1] = find(grid, memo, m, n+1)
	}
	return memo[m+1][n] + memo[m][n+1]
}

func main() {
	grid := [][]int{
		[]int{0, 0, 0},
		[]int{0, 1, 0},
		[]int{0, 0, 0},
	}
	fmt.Println(uniquePathsWithObstacles(grid))
}
