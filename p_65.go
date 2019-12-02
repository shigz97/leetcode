package main

import "fmt"

import "strings"

func isNumber(s string) bool {
	if len(s) == 0 || s == " " {
		return false
	}
	if s[len(s)-1] == ' ' {
		return isNumber(s[:len(s)-1])
	}
	switch strings.Count(s, "e") {
	case 1:
		i := strings.Index(s, "e")
		return isLeftVaild(s[:i]) && isRightValid(s[i+1:])
	case 0:
		return isLeftVaild(s)
	default:
		return false
	}
}

func isLeftVaild(s string) bool {
	if len(s) == 0 {
		return false
	}
	if s[0] == ' ' {
		return isLeftVaild(s[1:])
	}
	switch strings.Count(s, ".") {
	case 1:
		if s == "." {
			return false
		}
		i := strings.Index(s, ".")
		if i == len(s)-1 {
			return isNum(s[:len(s)-1])
		}
		if i == 0 {
			return isNum(s[1:])
		}
		return isNum(s[:i]) && isNum(s[i+1:])
	case 0:
		return isNum(s)
	default:
		return false
	}
}

func isRightValid(s string) bool {
	if len(s) == 0 {
		return false
	}
	switch strings.Count(s, ".") {
	case 0:
		return isNum(s)
	default:
		return false
	}
}

func isNum(s string) bool {
	if len(s) == 1 && (s == "+" || s == "-") {
		return false
	}
	if !(s[0] == '+' || s[0] == '-' || (s[0] >= '0' && s[0] <= '9')) {
		return false
	}
	for _, ss := range s[1:] {
		if !(ss >= '0' && ss <= '9') {
			return false
		}
	}
	return true
}

func main() {
	fmt.Println(isNumber(" "))
}
