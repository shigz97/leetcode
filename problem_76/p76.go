package main

import (
	"fmt"
	"strings"
)

func minWindow(s string, t string) string {
	if judge(s, t) == false {
		return ""
	}
	l, r, rl, rr := 0, 0, 0, 0
	for r+1 < len(s) {
		for judge(s[l:r+1], t) == false {
			r++
			if r >= len(s) {
				return s[rl : rr+1]
			}
		}
		for judge(s[l:r+1], t) == true {
			l++
		}

		if rr-rl == 0 || r-l+1 < rr-rl {
			rl, rr = l-1, r
		}
	}
	return s[rl : rr+1]
}

func judge(s string, t string) bool {
	res := true
	ishown := make(map[rune]bool, len(t))
	for _, i := range t {
		if ishown[i] == false {
			cs, ct := strings.Count(s, string(i)), strings.Count(t, string(i))
			if cs < ct {
				return false
			}
			ishown[i] = true
		}
	}
	return res
}

func main() {
	fmt.Println(minWindow("ab", "a"))
}
