package main

import "fmt"

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

/*type queue struct {
	node  []*TreeNode
	front int
	back  int
}*/

func createTree(nums []int) *TreeNode {
	if len(nums) == 0 || nums[0] == -1 {
		return nil
	}
	//root := &TreeNode{nums[0], nil, nil}
	/*q := queue{make([]*TreeNode, len(nums)), -1, -1}
	q.back++
	q.node[q.back] = root*/
	nodes := make([]*TreeNode, len(nums))
	for i := range nums {
		if nums[i] != -1 {
			nodes[i] = &TreeNode{nums[i], nil, nil}
			if i != 0 {
				if i%2 == 1 {
					nodes[i/2].Left = nodes[i]
				} else {
					nodes[i/2-1].Right = nodes[i]
				}
			}
		}
	}
	return nodes[0]
}

func getSeq(root *TreeNode, s *[]int) {
	if root != nil {
		getSeq(root.Left, s)
		*s = append(*s, root.Val)
		getSeq(root.Right, s)
	}
}

func isSymmetric(root *TreeNode) bool {
	seq := []int{}
	getSeq(root, &seq)
	center := len(seq) / 2
	for i := 0; i < center; i++ {
		if seq[i] != seq[center*2-i] {
			return false
		}
	}
	return true
}

func main() {
	nums := []int{1, 2, 2, 3, 4, 4, 3}
	root := createTree(nums)
	fmt.Println(isSymmetric(root))
}
