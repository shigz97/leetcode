package main

import "fmt"

func main() {
	s := "kkkkkk"
	fmt.Printf("%s\n", longestPalindrome(s))
}

func longestPalindrome(s string) string {
	r := reverse(s)
	var maxs string
	for i := 0; i < len(s); i = i + 1 {
		j := len(s) - i
		for k := len(s); k > i; k = k - 1 {
			p := len(s) - k
			if s[i:k] == r[p:j] {
				if k-i > len(maxs) {
					maxs = s[i:k]
				}
			}
		}
	}
	return maxs
}

func reverse(s string) string {
	r := []byte(s)
	for i, j := 0, len(r)-1; i < j; i, j = i+1, j-1 {
		r[i], r[j] = r[j], r[i]
	}
	return string(r)
}
