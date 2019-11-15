package main

import (
	"fmt"

	"github.com/shigzz/gokit/stack"
	. "github.com/shigzz/gokit/tree"
)

func kthSmallest(root *TreeNode, k int) int {
	s, _ := stack.NewStack("*TreeNode")
	n := 0
	for !s.IsEmpty() || root != nil {
		if root != nil {
			s.Push(root)
		} else {
			r, _ := s.Pop()
			root = r.(*TreeNode)
			n = n + 1
			if n == k {

			}
		}
	}
}

func main() {
	root := CreateTreeByInterface([]interface{}{5, 3, 6, 2, 4, "null", "null", 1})
	fmt.Println(kthSmallest(root, 3))
}
