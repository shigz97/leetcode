package main

import (
	"fmt"
	"github.com/shigzz/gokit/tree"
)

func pathSum(root *tree.TreeNode, sum int) [][]int {
	cura := []int{}
	res := [][]int{}
	if root == nil {
		return res
	}
	helper(root, sum, 0, &cura, &res)
	return res
}

func helper(root *tree.TreeNode, sum int, cur int, cura *[]int, res *[][]int) {
	cur = cur + root.Val
	*cura = append(*cura, root.Val)
	//fmt.Print(*cura)
	//fmt.Printf("%p\n",cura)
	if root.Left == nil && root.Right == nil {
		if cur == sum {
			new := append([]int{}, (*cura)...)
			*res = append(*res, new)
		}
	}
	if cur >= sum {
		//return
	}
	if root.Left != nil {
		helper(root.Left, sum, cur, cura, res)
	}
	if root.Right != nil {
		helper(root.Right, sum, cur, cura, res)
	}
	*cura = (*cura)[:len(*cura)-1]
}

func main() {
	root := tree.CreateTreeByInterface([]interface{}{5, 4, 8, 11, "null", 13, 4, 7, 2, "null", "null", 5, 1})
	tree.PrintTree(root)
	fmt.Println(pathSum(root, 22))
}
