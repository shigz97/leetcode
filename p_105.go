package main

import (
	"fmt"
	"leetcode/tree"
	. "leetcode/tree"
)

/*type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}*/

func buildTree(preorder []int, inorder []int) *TreeNode {
	if len(preorder) == 0 {
		return nil
	}
	tmp := preorder[0]
	loc := findloc(inorder, tmp)
	root := &TreeNode{tmp, nil, nil}
	root.Left = buildTree(preorder[1:loc+1], inorder[0:loc])
	root.Right = buildTree(preorder[loc+1:], inorder[loc+1:])
	return root
}
func findloc(nums []int, v int) int {
	for i, n := range nums {
		if n == v {
			return i
		}
	}
	return 0
}

func main() {
	pre := []int{3, 9, 20, 15, 7}
	in := []int{9, 3, 15, 20, 7}
	root := buildTree(pre, in)
	tree.PreOrderTree(root)
	fmt.Println(root)
}
