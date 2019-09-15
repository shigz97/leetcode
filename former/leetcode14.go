package main

import "fmt"

func longestCommonPrefix(strs []string) string {
	res := strs[0]
	//fmt.Println(res)
	for _, str := range strs[1:] {
		res = findsame(res, str)
	}
	fmt.Println(res)
	return res
}

func findsame(s, p string) string {
	i := 0
	for i < len(s) && i < len(p) {
		if s[i] != p[i] {
			break
		} else {
			i++
		}
	}
	return s[:i]
}

func main() {
	var strs = []string{"flower", "flow", "flight"}
	fmt.Println(longestCommonPrefix(strs))
}
