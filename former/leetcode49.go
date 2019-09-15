package main

import (
	"fmt"
	"sort"
)

func groupAnagrams(strs []string) [][]string {
	res := [][]string{}
	if len(strs) == 0 {
		return res
	}
	dic := make(map[string][]string, len(strs))
	for _, str := range strs {
		tmp := getsort(str)
		if _, ok := dic[tmp]; ok {
			dic[tmp] = append(dic[tmp], str)
		} else {
			dic[tmp] = append(dic[tmp], str)
		}
	}
	//fmt.Println(dic)
	for _, value := range dic {
		res = append(res, value)
	}
	return res
}

func getsort(s string) string {
	b := []byte(s)
	tmp := []int{}
	res := []byte{}
	for _, k := range b {
		tmp = append(tmp, int(k))
	}
	sort.Ints(tmp)
	for _, i := range tmp {
		res = append(res, byte(i))
	}
	return string(res)
}

func main() {
	strs := []string{
		"eat",
		"tea",
		"tan",
		"ate",
		"nat",
		"bat",
	}
	fmt.Println(groupAnagrams(strs))
}
