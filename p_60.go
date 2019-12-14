package main

import "fmt"

func getPermutation(n int, k int) string {
	ret := make([]byte, n)
	isVisit := make([]int, n+1)
	for i := 0; i < n; i++ {
		div := fact(n - i - 1)
		var index, remain int
		if k%div == 0 {
			index = k / div
			remain = div
		} else {
			index = k/div + 1
			remain = k % div
		}
		//fmt.Println(index, remain, 111)
		var is int
		sum := 0
		for j := 1; j <= n; j++ {
			if isVisit[j] == 1 {
				continue
			}
			sum++
			if sum == index {
				is = j
				break
			}
		}
		isVisit[is] = 1
		//fmt.Println(is, isVisit, 222)
		ret[i] = byte(is + 48)
		k = remain
	}
	return string(ret)
}

func fact(n int) int {
	if n == 1 || n == 0 {
		return 1
	}
	return n * fact(n-1)
}

func main() {
	fmt.Println(getPermutation(1, 1))
}
