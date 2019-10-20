package main

import (
	"fmt"
	"strings"
)

func removeSubfolders(folder []string) []string {
	res := []string{}
	for _, f := range folder {
		tmp := true
		var r string
		for i := 0; i < len(res); i++ {
			r = res[i]
			if strings.Contains(f, r) == true && strings.Count(f, "/") != strings.Count(r, "/") {
				tmp = false
				break
			}
			if strings.Contains(r, f) == true && strings.Count(f, "/") != strings.Count(r, "/") {
				res = append(res[:i], res[i+1:]...)
				i--
			}
		}
		if tmp == true {
			res = append(res, f)
		}
	}
	return res
}

func main() {
	//f := []string{"/a", "/a/b", "/c/d", "/c/d/e", "/c/f"}
	//f := []string{"/a/b/c", "/a/b/ca", "/a/b/d"}
	f := []string{"/aa/ab/ac/ae", "/aa/ab/af/ag", "/ap/aq/ar/as", "/ap/aq/ar", "/ap/ax/ay/az", "/ap", "/ap/aq/ar/at", "/aa/ab/af/ah", "/aa/ai/aj/ak", "/aa/ai/am/ao"}
	fmt.Println(removeSubfolders(f))
}
