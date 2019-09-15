package main

import "fmt"

type ListNode struct {
	Val  int
	Next *ListNode
}

func mergeTwoLists(l1 *ListNode, l2 *ListNode) *ListNode {
	if l1 == nil {
		return l2
	}
	if l2 == nil {
		return l1
	}
	head := &ListNode{}
	if l1.Val < l2.Val {
		head.Val = l1.Val
		l1 = l1.Next
	} else {
		head.Val = l2.Val
		l2 = l2.Next
	}
	head.Next = mergeTwoLists(l1, l2)
	return head
}

func main() {
	l1 := &ListNode{1, nil}
	l1.Next = &ListNode{2, nil}
	l1.Next.Next = &ListNode{4, nil}
	l2 := &ListNode{1, nil}
	l2.Next = &ListNode{3, nil}
	l2.Next.Next = &ListNode{4, nil}
	head := mergeTwoLists(l1, l2)
	for head != nil {
		fmt.Println(head.Val)
		head = head.Next
	}
}
