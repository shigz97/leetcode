package main

import (
	"fmt"
	. "leetcode/tree"
)

func lowestCommonAncestor(root, p, q *TreeNode) *TreeNode {
	_, res := findAncestor(root, p, q)
	return res
}

func findAncestor(root, p, q *TreeNode) (int, *TreeNode) {
	if root == nil {
		return 0, nil
	}
	var center, left, right int
	var resl, resr *TreeNode
	left, resl = findAncestor(root.Left, p, q)
	if resl != nil {
		return 1, resl
	}
	right, resr = findAncestor(root.Right, p, q)
	if resr != nil {
		return 1, resr
	}
	if root == p || root == q {
		center = 1
	}
	if left == 0 && right == 0 && center == 0 {
		return 0, nil
	} else if left^right^center == 0 {
		return 1, root
	} else {
		return 1, nil
	}
}

func main() {
	objs := []interface{}{3, 5, 1, 6, 2, 0, 8, "null", "null", 7, 4}
	root := CreateTreeByInterface(objs)
	node := lowestCommonAncestor(root, root.Left, root.Left.Right)
	fmt.Println(node)
}
