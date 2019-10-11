package main

import (
	. "leetcode/tree"
)

func invertTree(root *TreeNode) *TreeNode {
	if root == nil {
		return nil
	}
	if root.Left != nil {
		invertTree(root.Left)
	}
	if root.Right != nil {
		invertTree(root.Right)
	}
	root.Left, root.Right = root.Right, root.Left
	return root
}

func main() {
	nums := []int{4, 2, 7, 1, 3, 6, 9}
	head := CreateTree(nums)
	head = invertTree(head)
	PrintTree(head)
}
