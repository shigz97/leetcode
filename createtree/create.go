package createtree

type Treenode struct {
	Val   int
	Left  *Treenode
	Right *Treenode
}

/*type queue struct {
	node  []*Treenode
	front int
	back  int
}*/

func createTree(nums []int) *Treenode {
	if len(nums) == 0 || nums[0] == -1 {
		return nil
	}
	//root := &Treenode{nums[0], nil, nil}
	/*q := queue{make([]*Treenode, len(nums)), -1, -1}
	q.back++
	q.node[q.back] = root*/
	nodes := make([]*Treenode, len(nums))
	for i := range nums {
		if nums[i] != -1 {
			nodes[i] = &Treenode{nums[i], nil, nil}
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
