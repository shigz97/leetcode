package main

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func isValidBST(root *TreeNode) bool {
	return judge(root, 0, 0)
}

func judge(root *TreeNode, val int, sign int) bool {
	res := true
	if root.Left == nil && root.Right == nil {
		return res
	}
	if sign == 0 {
		if root.Left != nil {
			if root.Val < root.Left.Val {
				res = false
			} else {
				res = res && judge(root.Left, root.Val, 1)
			}
		}
		if root.Right != nil {
			if root.Val > root.Right.Val {
				res = false
			} else {
				res = res && judge(root.Right, root.Val, 2)
			}
		}
		return res
	} else if sign == 1 {
		if root.Left != nil {
			if root.Val < root.Left.Val {
				res = false
			} else {
				res = res && judge(root.Left, root.Val, 1)
			}
		}
		if root.Right != nil {
			if root.Val > root.Right.Val {
				res = false
			} else if val < root.Right.Val {
				res = false
			} else {
				res = res && judge(root.Right, root.Val, 2)
			}
		}
		return res
	} else {
		if root.Left != nil {
			if root.Val < root.Left.Val {
				res = false
			} else if val > root.Left.Val {
				res = false
			} else {
				res = res && judge(root.Left, root.Val, 1)
			}
		}
		if root.Right != nil {
			if root.Val > root.Right.Val {
				res = false
			} else {
				res = res && judge(root.Right, root.Val, 2)
			}
		}
	}
}
