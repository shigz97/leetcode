package main

import (
	. "github.com/shigzz/gokit/tree"
)

func sortedArrayToBST(nums []int) *TreeNode {
	if len(nums) == 0 {
		return nil
	}
	c := len(nums) / 2
	root := &TreeNode{nums[c], nil, nil}
	root.Left = sortedArrayToBST(nums[:c])
	root.Right = sortedArrayToBST(nums[c+1:])
	return root
}

func main() {
	nums := []int{-10, -3, 0, 5, 9}
	root := sortedArrayToBST(nums)
	PrintTree(root)
}
