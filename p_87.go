package main

import "fmt"

import "sort"

func isScramble(s1 string, s2 string) bool {
	if s1 == s2 {
		return true
	}
	b1, b2 := []byte(s1), []byte(s2)
	sort.Slice(b1, func(i, j int) bool {
		return b1[i] < b1[j]
	})
	sort.Slice(b2, func(i, j int) bool {
		return b2[i] < b2[j]
	})
	if string(b1) != string(b2) {
		return false
	}
	for i := 1; i < len(s1); i++ {
		if isScramble(s1[:i], s2[:i]) && isScramble(s1[i:], s2[i:]) {
			return true
		}
		if isScramble(s1[:i], s2[len(s1)-i:]) && isScramble(s1[i:], s2[:len(s1)-i]) {
			return true
		}
	}
	return false
}

func main() {
	fmt.Println(isScramble("eat", "ate"))
}
