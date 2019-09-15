package main

import "fmt"

//自顶向下
/*func uniquePaths(m int, n int) int {
	memo := [][]int{}
	for i := 0; i < m; i++ {
		tmp := make([]int, n)
		memo = append(memo, tmp)
	}
	return find(m-1, n-1, memo)
	return 0
}

func find(m, n int, memo [][]int) int {
	if m == 0 || n == 0 {
		return 1
	}
	if memo[m-1][n] == 0 {
		memo[m-1][n] = find(m-1, n, memo)
	}
	if memo[m][n-1] == 0 {
		memo[m][n-1] = find(m, n-1, memo)
	}
	return memo[m-1][n] + memo[m][n-1]
}*/

func uniquePaths(m int, n int) int {
	memo := [][]int{}
	for i := 0; i < m; i++ {
		tmp := make([]int, n)
		memo = append(memo, tmp)
	}
	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			if i == 0 || j == 0 {
				memo[i][j] = 1
			} else {
				memo[i][j] = memo[i-1][j] + memo[i][j-1]
			}
		}
	}
	return memo[m-1][n-1]
}

func main() {
	fmt.Println(uniquePaths(7, 3))
}
