package main

import "fmt"

import "strconv"

func numDecodings(s string) int {
	if s == "0" {
		return 0
	}
	if len(s) == 0 || len(s) == 1 {
		return len(s)
	}
	if len(s) == 2 {
		tmp, _ := strconv.Atoi(s)
		if tmp <= 26 {
			if tmp%10 != 0 {
				return 2
			}
			return 1
		}
		if s[1] != '0' {
			return 1
		}
		return 0
	}
	var res = 0
	if s[1] != '0' {
		res += numDecodings(s[1:])
	}
	tp, _ := strconv.Atoi(s[:2])
	if tp <= 26 {
		res += numDecodings(s[2:])
	}
	return res
}

func main() {
	fmt.Println(numDecodings("10"))
}
