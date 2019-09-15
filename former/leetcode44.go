package main

import "fmt"

func isMatch(s string, p string) bool {
	fmt.Println(s, p)
	if len(p) == 0 && len(s) == 0 {
		return true
	} else {
		if len(p) == 0 {
			return false
		}
	}
	if p[0] == '*' {
		for len(p) > 1 && p[1] == '*' {
			p = p[1:]
		}
		if len(s) == 0 {
			return isMatch(s, p[1:])
		} else {
			return isMatch(s[1:], p) || isMatch(s, p[1:])
		}
	}
	if len(s) == 0 || (p[0] != '?' && s[0] != p[0]) {
		return false
	} else {
		return isMatch(s[1:], p[1:])
	}
}

func main() {
	s, p := "bbbbbbbabbaabbabbbbaaabbabbabaaabbababbbabbbabaaabaab", "b*b*ab**ba*b**b***bba"
	fmt.Println(isMatch(s, p))
}
