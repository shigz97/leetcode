package main

import "fmt"

func mySqrt(x int) int {
	res := 1
	for res*res <= x {
		res++
	}
	return res - 1
}

func main() {
	fmt.Println(mySqrt(3))
}
