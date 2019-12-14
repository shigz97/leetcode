package main

import "fmt"

import "strings"

func simplifyPath(path string) string {
	p := strings.Split(path, "/")
	ret := make([]string, len(p))
	ret[0] = ""
	top := 0
	for _, pp := range p {
		switch {
		case pp == "" || pp == ".":
			continue
		case pp == "..":
			if top > 0 {
				top--
			}
		default:
			top++
			ret[top] = pp
		}
	}
	res := ""
	for i := 0; i < top; i++ {
		res = res + ret[i] + "/"
	}
	res = res + ret[top]
	if top == 0 {
		res = "/"
	}
	return res
}

func main() {
	fmt.Println(simplifyPath("/a//b////c/d//././/.."))
}
