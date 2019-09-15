package main

import "fmt"

func minDistance(word1 string, word2 string) int {
	m, n := len(word1), len(word2)
	if m == 0 {
		return n
	}
	if n == 0 {
		return m
	}
	res := make([][]int, m+1)
	for i := range res {
		res[i] = make([]int, n+1)
	}
	for i := 0; i <= n; i++ {
		res[0][i] = i
	}
	for i := 0; i <= m; i++ {
		res[i][0] = i
	}
	for i := 1; i <= m; i++ {
		for j := 1; j <= n; j++ {
			if word1[i-1] == word2[j-1] {
				res[i][j] = 1 + getMin(res[i-1][j-1]-1, res[i][j-1], res[i-1][j])
			} else {
				res[i][j] = 1 + getMin(res[i-1][j-1], res[i][j-1], res[i-1][j])
			}
		}
	}
	return res[m][n]
}

func getMin(i1, i2, i3 int) int {
	min := i1
	if i2 < min {
		min = i2
	}
	if i3 < min {
		min = i3
	}
	return min
}

func main() {
	fmt.Println(minDistance("horse", "ros"))
}
