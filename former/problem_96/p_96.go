package main

import "fmt"

func numTrees(n int) int {
	if n <= 0 {
		return 0
	}
	memo := make([]int, n+1)
	memo[1] = 1
	for i := 2; i <= n; i++ {
		tmp := memo[i-1]
		for j := 1; j <= i/2-1; j++ {
			tmp = tmp + memo[j]*memo[i-1-j]
		}
		if i%2 == 1 {
			memo[i] = 2*tmp + memo[i/2]*memo[i/2]
		} else {
			memo[i] = 2 * tmp
		}
	}
	return memo[n]
}

func main() {
	fmt.Println(numTrees(6))
}
