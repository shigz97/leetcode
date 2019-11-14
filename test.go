package main

import (
	"fmt"
	. "tree"
)

func main() {
	head := CreateTree([]int{1, 2, 3, 4, 5, 6})
	fmt.Println(head)
	PrintTree(head)
}
