package main

import (
	"fmt"
	. "leetcode/tree"
	"math"
)

func maxPathSum(root *TreeNode) int {
	max := math.MinInt64
	findMax(root, &max)
	return max
}

func findMax(root *TreeNode, max *int) int {
	if root != nil {
		ml := findMax(root.Left, max)
		mr := findMax(root.Right, max)
		res := root.Val
		if ml > 0 {
			res = res + ml
		}
		if mr > 0 {
			res = res + mr
		}
		if res > *max {
			*max = res
		}
		return res
	} else {
		return 0
	}
}

func main() {
	root := CreateTreeByInterface(5, 4, 8, 11, "null", 13, 4, 7, 2, "null", "null", "null", 1)
	fmt.Println(maxPathSum(root))
}
