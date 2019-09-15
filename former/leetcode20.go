package main

import "fmt"

func isVaild(s string) bool {
	cha := make([]byte, len(s))
	top := -1
	for i := 0; i < len(s); i++ {
		if s[i] == '(' || s[i] == '{' || s[i] == '[' {
			top++
			cha[top] = s[i]
		} else {
			if top == -1 {
				return false
			} else {
				tmp := cha[top]
				top--
				switch {
				case s[i] == ')':
					if tmp != '(' {
						return false
					}
				case s[i] == '}':
					if tmp != '{' {
						return false
					}
				case s[i] == ']':
					if tmp != '[' {
						return false
					}
				}
			}
		}
	}
	if top == -1 {
		return true
	} else {
		return false
	}
}

func main() {
	fmt.Println(isVaild("(){}[]"))
}
