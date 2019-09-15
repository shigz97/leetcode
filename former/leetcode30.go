package main

import (
	"fmt"
	"sort"
)

func findSubstring(s string, words []string) []int {
	res := []int{}
	dic := make(map[int]string, len(s))
	pos := make([]int, len(words))
	for i, str := range words {
		p := getPos(s, str)
		if k, ok := dic[p]; ok {
			p = getPos(s[p+len(k):], str) + p + len(k)
			dic[p] = str
		} else {
			dic[p] = str
		}
		if p == -1 {
			return res
		}
		pos[i] = p
	}
	ii := 0
	for ii < len(s) {
		sort.Ints(pos)
		ii = pos[0]
		flag := 1
		for i := 0; i < len(pos)-1; i++ {
			if pos[i]+len(dic[pos[i]]) == pos[i+1] {
				continue
			} else {
				flag = 0
				break
			}
		}
		if flag == 1 {
			res = append(res, ii)
		}
		tmp := pos[0]
		pp := getPos(s[tmp+len(dic[tmp]):], dic[tmp]) + tmp + len(dic[tmp])
		if k, ok := dic[pp]; ok {
			pp = getPos(s[pp+len(k):], k) + pp + len(k)
			dic[pp] = k
		} else {
			dic[pp] = k
		}
		if pp == -1 {
			return res
		} else {
			pos[0] = pp
		}
	}
	return res
}

func getPos(s, p string) int {
	next := getNext(p)
	i, j := 0, 0
	for i < len(s) && j < len(p) {
		if s[i] == p[j] {
			i, j = i+1, j+1
		} else {
			j = next[j]
			if j == -1 {
				i, j = i+1, 0
			}
		}
	}
	if j == len(p) {
		return i - len(p)
	}
	return -1
}

func getNext(p string) []int {
	i, j := 0, -1
	next := make([]int, len(p))
	next[0] = -1
	for i < len(p)-1 {
		if j == -1 || p[i] == p[j] {
			i, j = i+1, j+1
			next[i] = j
		} else {
			j = next[j]
		}
	}
	return next
}

func main() {
	//fmt.Println(getNext("abcabaabcd"))
	s := "wordgoodgoodgoodbestword"
	words := []string{"word", "good", "best", "word"}
	fmt.Println(findSubstring(s, words))
}
