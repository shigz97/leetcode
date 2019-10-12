package tree

func CreateTreeByInterface(nums []interface{}) *TreeNode {
	if len(nums) == 0 || nums[0] == "null" {
		return nil
	}
	nodes := make([]*TreeNode, len(nums))
	for i := range nums {
		if nums[i] != "null" {
			//nodes[i] = &TreeNode{nums[i], nil, nil}
			n, _ := nums[i].(int)
			nodes[i] = &TreeNode{n, nil, nil}
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
