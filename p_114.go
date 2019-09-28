package leetcode

func flatten(root *TreeNode) {
	if root != nil {
		flatten(root.Left)
		flatten(root.Right)
		if root.Left != nil {
			tp := root.Left
			for tp.Right != nil {
				tp = tp.Right
			}
			tp.Right = root.Right
			root.Right = root.Left
			root.Left = nil
		}
	}
	return
}
