package tree

import "fmt"

func PreOrderTree(root *TreeNode) {
	if root != nil {
		fmt.Print(root.Val, " ")
		PreOrderTree(root.Left)
		PreOrderTree(root.Right)
	}
}

func InOrderTree(root *TreeNode) {
	if root != nil {
		InOrderTree(root.Left)
		fmt.Print(root.Val, " ")
		InOrderTree(root.Right)
	}
}

func PostOrderTree(root *TreeNode) {
	if root != nil {
		PostOrderTree(root.Left)
		PostOrderTree(root.Right)
		fmt.Print(root.Val, " ")
	}
}

func PrintTree(root *TreeNode) {
	fmt.Print("PreOrder: ")
	PreOrderTree(root)
	fmt.Println()
	fmt.Print("InOrder: ")
	InOrderTree(root)
	fmt.Println()
	fmt.Print("PostOrder: ")
	PostOrderTree(root)
	fmt.Println()
}
