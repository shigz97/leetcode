package leetcode

func levelOrder(root *TreeNode) [][]int {
	res := [][]int{}
	if root == nil {
		return res
	}
	nodes := []*TreeNode{}
	front, rear := -1, -1
	rear++
	nodes = append(nodes, root)
	back := root
	tmp := []int{}
	for rear > front {
		front++
		tp := nodes[front]
		tmp = append(tmp, tp.Val)
		if tp.Left != nil {
			rear++
			nodes = append(nodes, tp.Left)
		}
		if tp.Right != nil {
			rear++
			nodes = append(nodes, tp.Right)
		}
		if tp == back {
			t := append([]int{}, tmp...)
			res = append(res, t)
			tmp = []int{}
			back = nodes[rear]
		}
	}
	return res
}
