package main

import "fmt"

func fullJustify(words []string, maxWidth int) []string {
	res := []string{}
	findj(&words, 0, maxWidth, &res)
	return res
}

func findj(word *[]string, i int, maxWidth int, res *[]string) {
	if i == len(*word) {
		return
	}
	w := *word
	nowl, length, tmp := 0, 0, []string{}
	for i < len(w) && nowl+len(w[i]) <= maxWidth {
		tmp = append(tmp, w[i])
		nowl = nowl + len(w[i]) + 1
		i++
		length++
	}
	//fmt.Println(tmp, length)
	ret := format(tmp, length, maxWidth)
	//fmt.Println(ret, len(ret))
	*res = append(*res, ret)
	findj(word, i, maxWidth, res)
}

func format(str []string, length int, maxWidth int) string {
	ret := ""
	d := make([]int, length-1)
	delta := maxWidth
	for _, s := range str {
		delta -= len(s)
	}
	ret = ret + str[0]
	//fmt.Println(ret)
	if length == 1 {
		ret = ret + formatstring(delta)
		return ret
	}
	for i := 0; i < length-1; i++ {
		d[i] = delta / (length - 1)
	}
	for i := 0; i < delta%(length-1); i++ {
		d[i]++
	}
	for i := 0; i < length-1; i++ {
		ret = ret + formatstring(d[i])
		ret = ret + str[i+1]
		//fmt.Println(ret)
	}
	return ret
}

func formatstring(n int) string {
	ss := make([]rune, n)
	for i := range ss {
		ss[i] = ' '
	}
	return string(ss)
}

func main() {
	//words := []string{"This", "is", "an", "example", "of", "text", "justification."}
	words := []string{"What", "must", "be", "acknowledgment", "shall", "be"}
	fmt.Println(words)
	fmt.Println(fullJustify(words, 16))
}
