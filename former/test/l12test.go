package main

import "fmt"

func isMatch(s string, p string) bool {
	if p == "" {
		return s == ""
	}
	var first bool
	if p != "" && s != "" {
		first = s[0] == p[0] || p[0] == '.'
	}
	fmt.Println(s, p, first)
	if len(p) >= 2 && p[1] == '*' {
		return isMatch(s, p[2:]) || (first && isMatch(s[1:], p))
	} else {
		return first && isMatch(s[1:], p[1:])
	}
	return false
}

func main() {
	s, p := "a", "a*"
	fmt.Println(isMatch(s, p))
}
