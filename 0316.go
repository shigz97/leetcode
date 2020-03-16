package main

import (
	"fmt"
	"strconv"
)

func main() {
	str := "aabbcccddwaaa"
	fmt.Println(compressString(str))
}

func compressString(S string) string {
	b := []byte(S)
	i := 0
	//res := []byte{}
	res := ""
	for i < len(S) {
		//res = append(res, b[i])
		res = res + string(b[i])
		sum, j := 1, i+1
		for j < len(S) && b[j] == b[i] {
			sum, j = sum+1, j+1
		}
		//fmt.Println(i, j, sum)
		//res = append(res, byte(sum+48))
		res = res + strconv.Itoa(sum)
		i = j
	}
	if len(res) > len(b) {
		return S
	}
	return string(res)
}
