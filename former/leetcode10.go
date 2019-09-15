package main

import "fmt"

func isMatch(s string, p string) bool {
	i, j := 0, 0
	var tmp byte
	if len(p) == 0 {
		if len(s) == 0 {
			return true
		} else {
			return false
		}
	} else {
		if len(s) == 0 {
			return false
		}
	}

	if p[i] == '.' || p[i] == s[j] {
		tmp = p[i]
	} else {
		return false
	}
	if i+1 < len(p) && p[i+1] == '*' {
		if p[i] == '.' {
			i = i + 2
			j = len(s) - len(p) + i
			return isMatch(s[j:], p[i:])
		} else {
			//fmt.Println(s, p, i, j)
			for s[j] == tmp && j+1 < len(s) {
				j = j + 1
			}
			if i+2 < len(p) {
				i = i + 2
			}
			if i+2 == len(p) && j+1 == len(s) {
				return true
			}
			//fmt.Println(s, p, i, j)
			return isMatch(s[j:], p[i:])
		}
	} else {
		i, j = i+1, j+1
		fmt.Println(s, p, i, j)
		return isMatch(s[j:], p[i:])
	}

}

func main() {
	s, p := "", ""
	fmt.Println(isMatch(s, p))
}
