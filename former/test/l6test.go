package main

import "fmt"

func main() {
	s := "PAYPALISHIRING"
	fmt.Println(convert(s, 3))
}

func convert(s string, numRows int) string {
	dic := make(map[int][]int, numRows)
	mult := 2*numRows - 2
	var j int
	for i := 1; i <= len(s); i++ {
		j = i % mult
		if j == 0 {
			j = mult
		}
		if j > numRows {
			j = 2*numRows - j
		}
		dic[j] = append(dic[j], i)
	}
	var r []byte
	for i := 1; i <= numRows; i++ {
		for _, k := range dic[i] {
			//fmt.Println(k)
			r = append(r, s[k-1])
		}
	}
	fmt.Println(dic)
	return string(r)
}
