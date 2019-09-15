package main

import "fmt"

func letterCombinations(digits string) []string {
	var dic = map[byte][]string{
		'2': []string{"a", "b", "c"},
		'3': []string{"d", "e", "f"},
		'4': []string{"g", "h", "l"},
		'5': []string{"j", "k", "l"},
		'6': []string{"m", "n", "o"},
		'7': []string{"p", "q", "r", "s"},
		'8': []string{"t", "u", "v"},
		'9': []string{"w", "x", "y", "z"},
	}
	if len(digits) == 0 {
		return []string{}
	}
	res := dic[digits[0]]
	for _, b := range digits[1:] {
		res = add(dic, res, byte(b))
	}
	return res
}

func add(dic map[byte][]string, former []string, x byte) []string {
	r := make([]string, 0)
	for _, a := range former {
		for _, b := range dic[x] {
			r = append(r, a+b)
			fmt.Println(r)
		}
	}
	return r
}

func main() {
	d := "233"
	fmt.Println(letterCombinations(d))
}
