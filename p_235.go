package main

import (
	"fmt"

	. "github.com/shigzz/gokit/tree"
)

func lowestCommonAncestor(root, p, q *TreeNode) *TreeNode {
	switch {
	case root.Val > p.Val && root.Val > q.Val:
		return lowestCommonAncestor(root.Left, p, q)
	case root.Val < p.Val && root.Val < q.Val:
		return lowestCommonAncestor(root.Right, p, q)
	default:
		return root
	}
}

func main() {
	root := CreateTreeByInterface([]interface{}{6, 2, 8, 0, 4, 7, 9, "null", "null", 3, 5})
	node := lowestCommonAncestor(root, root.Left, root.Left.Right)
	fmt.Println(node)
}
