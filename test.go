package main

import (
	. "leetcode/tree"
)

func main() {
	root := CreateTreeByInterface(1, "null", 2, "null", "null", "null", "null")
	PrintTree(root)
}
