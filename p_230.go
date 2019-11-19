package main

import (
	"fmt"

	. "github.com/shigzz/gokit/tree"
)

func kthSmallest(root *TreeNode, k int) int {
	s, top, num := []*TreeNode{}, -1, 0
	for top != -1 || root != nil {
		if root != nil {
			top = top + 1
			if top >= len(s) {
				s = append(s, root)
			} else {
				s[top] = root
			}
			root = root.Left
		} else {
			num = num + 1
			if num == k {
				return s[top].Val
			}
			root = s[top].Right
			top = top - 1
		}
	}
	return 0
}

func main() {
	root := CreateTreeByInterface([]interface{}{5, 3, 6, 2, 4, "null", "null", 1})
	fmt.Println(kthSmallest(root, 3))
}
