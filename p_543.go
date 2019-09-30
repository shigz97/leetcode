package main

import (
	"fmt"
	. "leetcode/tree"
)

func diameterOfBinaryTree(root *TreeNode) int {
	max := 0
	getHeight(root, &max)
	return max
}

func getHeight(root *TreeNode, max *int) int {
	if root != nil {
		hl := getHeight(root.Left, max)
		rl := getHeight(root.Right, max)
		if hl+rl > *max {
			*max = hl + rl
		}
		if hl > rl {
			return hl + 1
		} else {
			return rl + 1
		}
	} else {
		return 0
	}
}

func main() {
	root := CreateTree([]int{1, 2, 3, 4, 5})
	fmt.Println(diameterOfBinaryTree(root))
}
