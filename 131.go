package main

import "fmt"

func main() {
	fmt.Println(partition("aab"))
}

func partition(s string) [][]string {
	res := [][]string{}
	if len(s) == 0 {
		return res
	}
	dfs(s, []string{}, &res)
	return res
}

func dfs(s string, cur []string, res *[][]string) {
	if len(s) == 0 {
		*res = append(*res, cur)
		return
	}
	for i := 1; i <= len(s); i++ {
		if judge(s[:i]) {
			cur = append(cur, s[:i])
			dfs(s[i:], cur, res)
		}
		if len(cur) > 0 {
			cur = cur[:len(cur)-1]

		}
	}
}

func judge(s string) bool {
	b := []byte(s)
	for i, j := 0, len(b)-1; i < j; i, j = i+1, j-1 {
		if b[i] != b[j] {
			return false
		}
	}
	return true
}
