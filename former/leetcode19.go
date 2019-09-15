package main

import "fmt"

type ListNode struct {
	Val  int
	Next *ListNode
}

func removeNthFromEnd(head *ListNode, n int) *ListNode {
	res, tmp := head, head
	//pre := &ListNode{}
	ad := 0
	for tmp.Next != nil {
		tmp = tmp.Next
		ad++
		if ad > n {
			res = res.Next
		}
	}
	if ad+1 == n {
		return head.Next
	} else {
		res.Next = res.Next.Next
	}
	return head
}

func main() {
	head := &ListNode{1, nil}
	p := head
	for i := 2; i <= 5; i++ {
		tmp := &ListNode{i, nil}
		p.Next = tmp
		p = p.Next
	}
	head = removeNthFromEnd(head, 2)
	for head != nil {
		fmt.Println(head.Val)
		head = head.Next
	}
}
