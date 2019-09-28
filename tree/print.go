package tree

import "fmt"

func PreOrderTree(root *TreeNode) {
	if root != nil {
		fmt.Println(root.Val)
		PreOrderTree(root.Left)
		PreOrderTree(root.Right)
	}
}

func InOrderTree(root *TreeNode) {
	if root != nil {
		InOrderTree(root.Left)
		fmt.Println(root.Val)
		InOrderTree(root.Right)
	}
}

func PostOrderTree(root *TreeNode) {
	if root != nil {
		PostOrderTree(root.Left)
		PostOrderTree(root.Right)
		fmt.Println(root.Val)
	}
}
