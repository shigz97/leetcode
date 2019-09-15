package main

import "fmt"

func main() {
	i1 := []int{1, 2, 3, 4, 5, 6}
	for i := range i1 {
		i = i + 1
		fmt.Println(i1[i])
	}
}
