package main

import (
	"fmt"
	. "github.com/shigzz/gokit/tree"
)

func zigzagLevelOrder(root *TreeNode) [][]int {
	res := [][]int{}
	if root == nil {
		return res
	}
	var tmp = []int{}
	queue := []*TreeNode{}
	front, end, last := -1, 0, 0
	//queue[end] = root
	queue = append(queue, root)
	for front < end {
		front++
		tmpn := queue[front]
		tmp = append(tmp, tmpn.Val)
		if tmpn.Left != nil {
			end++
			//queue[end] = tmpn.Left
			queue = append(queue, tmpn.Left)
		}
		if tmpn.Right != nil {
			end++
			//queue[end] = tmpn.Right
			queue = append(queue, tmpn.Right)
		}
		if front == last {
			last = end
			res = append(res, tmp)
			tmp = make([]int, 0)
		}
	}
	for i := 1; i < len(res); i += 2 {
		reverse(res[i])
	}
	return res
}

func reverse(nums []int) {
	for i, j := 0, len(nums)-1; i < j; i, j = i+1, j-1 {
		nums[i], nums[j] = nums[j], nums[i]
	}
}

func main() {
	root := CreateTreeByInterface([]interface{}{3, 9, 20, "null", "null", 15, 7})
	fmt.Println(zigzagLevelOrder(root))
	//PrintTree(root)
}
