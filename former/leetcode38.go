package main

import (
	"fmt"
	"strconv"
)

func countAndSay(n int) string {
	if n == 1 {
		return "1"
	}
	str := countAndSay(n - 1)
	res := ""
	tmp, count := str[0], 1
	for i := 1; i < len(str); i++ {
		if str[i] == str[i-1] {
			count++
		} else {
			res = res + strconv.Itoa(count) + string(tmp)
			tmp, count = str[i], 1
		}
	}
	res = res + strconv.Itoa(count) + string(tmp)
	return res
}

func main() {
	fmt.Println(countAndSay(30))
}
