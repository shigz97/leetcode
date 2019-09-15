package main

import "fmt"

func main() {
	tmp := []int{1, 2, 3}
	res := [][]int{}
	res = append(res, tmp)
	fmt.Println(res)
	tmp[0] = 4
	fmt.Println(res)
}
