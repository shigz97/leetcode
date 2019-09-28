package tree

func CreateTreeByInterface(nums []interface{}) *Treenode {
	if len(nums) == 0 || nums[0] == "null" {
		return nil
	}
	nodes := make([]*Treenode, len(nums))
	for i := range nums {
		if nums[i] != "null" {
			//nodes[i] = &Treenode{nums[i], nil, nil}
			n, _ := nums[i].(int)
			nodes[i] = &Treenode{n, nil, nil}
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
