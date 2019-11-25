package main

import (
	. "github.com/shigzz/gokit/linklist"
)

var head = CreateList([]int{4, 5, 1, 9})

func deleteNode(node *ListNode) {
	if node == head {
		head = head.Next
		return
	}
	pre, now := head, head.Next
	for now != node {
		pre, now = pre.Next, now.Next
	}
	pre.Next = now.Next
	return
}

func main() {
	deleteNode(head.Next.Next)
	PrintList(head)
}
