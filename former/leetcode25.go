package main

import "fmt"

type ListNode struct {
	Val  int
	Next *ListNode
}

func reverseKGroup(head *ListNode, k int) *ListNode {
	if head == nil {
		return nil
	}
	var p1, p2 *ListNode
	p1, p2 = head, head
	for check(p1, k) == true {
		pre := p2
		p2 = p1
		for i := k; i > 1; i-- {
			p2 = p2.Next
		}
		t1 := reverse(p1, p2, k)
		if p1 == head {
			head = t1
		} else {
			pre.Next = t1
		}
		p2 = p1
		p1 = p1.Next
	}
	return head
}

func check(head *ListNode, k int) bool {
	for head != nil {
		k--
		if k == 0 {
			return true
		}
		head = head.Next
	}
	return false
}

func reverse(p1, p2 *ListNode, k int) *ListNode {
	for p1 != p2 {
		tmp := p1.Next
		p1.Next = p2.Next
		p2.Next = p1
		p1 = tmp
	}
	return p1
}

func main() {
	l1 := &ListNode{1, nil}
	l1.Next = &ListNode{2, nil}
	l1.Next.Next = &ListNode{3, nil}
	l1.Next.Next.Next = &ListNode{4, nil}
	l1.Next.Next.Next.Next = &ListNode{5, nil}
	head := reverseKGroup(l1, 3)
	for head != nil {
		fmt.Println(head.Val)
		head = head.Next
	}
}
