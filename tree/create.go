package tree

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

func CreateTree(nums []int) *TreeNode {
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
