package main

import "fmt"

func longestValidParentheses(s string) int {
	dp := make([]int, len(s))
	cp := make([]int, len(s))
	max := 0
	extra := 0
	if s[0] == '(' {
		dp[0], cp[0] = 0, 1
	} else {
		dp[0], cp[0] = 0, 0
	}
	for i := 1; i < len(s); i++ {
		if s[i] == '(' {
			dp[i], cp[i] = dp[i-1], cp[i-1]+1
		} else {
			cp[i] = cp[i-1] - 1
			if cp[i] < 0 {
				dp[i], cp[i] = 0, 0
			} else {
				dp[i] = dp[i-1] + 2
				if dp[i] > max {
					max = dp[i]
				}
			}
		}
		fmt.Println(i, dp[i], cp[i])
	}
	return max
}

func main() {
	kuohao := "()((()(())"
	fmt.Println(longestValidParentheses(kuohao))
}
