package main

import (
	"fmt"
	"leetcode/tree"
)

func main() {
	root := tree.CreateTree([]int{1, 2, 3, 4, 5, 6})
	fmt.Println(root.Left.Val)
}
