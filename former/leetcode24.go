package main

import "fmt"

type ListNode struct {
	Val  int
	Next *ListNode
}

func swapPairs(head *ListNode) *ListNode {
	p1, p2 := head, head.Next
	for p1 != nil && p1.Next != nil {
		if p1 == head {
			p1.Next = p2.Next
			p2.Next = p1
			head = p2
			p1, p2 = p1.Next, p2.Next
		} else {
			pre := p2
			p2 = p1.Next
			pre.Next = p2
			p1.Next = p2.Next
			p2.Next = p1
			p1, p2 = p1.Next, p2.Next
		}
	}
	return head
}

func main() {
	l1 := &ListNode{1, nil}
	l1.Next = &ListNode{2, nil}
	l1.Next.Next = &ListNode{3, nil}
	l1.Next.Next.Next = &ListNode{4, nil}
	l1.Next.Next.Next.Next = &ListNode{5, nil}
	head := swapPairs(l1)
	for head != nil {
		fmt.Println(head.Val)
		head = head.Next
	}
}
