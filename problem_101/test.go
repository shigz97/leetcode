package main

import (
	"fmt"
)

type Treenode struct {
	Val   int
	Left  *Treenode
	Right *Treenode
}

func main() {
	nums := []int{1, 2, 2, -1, 3, -1, 3}
	root := createTree(nums)
	print(root)
}

func print(root *Treenode) {
	if root != nil {

		print(root.Left)

		print(root.Right)
		fmt.Println(root.Val)
	}
}
