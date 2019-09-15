package main

import "fmt"

func main() {
	res := make([]int, 6)
	res[0], res[1], res[2] = 1, 1, 1

	for _, n := range res {
		fmt.Println(n)
	}
}
