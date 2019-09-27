func maxDepth(root *TreeNode) int {
	if root == nil {
		return 0
	}
	hl, h2 := maxDepth(root.Left), maxDepth(root.Right)
	if hl > h2 {
		return hl + 1
	} else {
		return h2 + 1
	}
}