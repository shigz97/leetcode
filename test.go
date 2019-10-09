package main

import "fmt"

func main() {
	i1 := []int{1, 2, 3}
	fmt.Println(&i1[0])
	i2 := append([]int{}, i1...)
	fmt.Println(&i2[0])
}
