package main

import (
	. "leetcode/linklist"
)

func sortList(head *ListNode) *ListNode {
	if head == nil {
		return nil
	}
	slow, fast := head, head.Next
	if fast == nil {
		return head
	}
	for fast != nil && fast.Next != nil {
		slow, fast = slow.Next, fast.Next.Next
	}
	tmp := slow.Next
	//fmt.Println(tmp)
	slow.Next = nil
	head = sortList(head)
	tmp = sortList(tmp)
	return mergeList(head, tmp)
}

func mergeList(head, center *ListNode) *ListNode {
	res := &ListNode{0, nil}
	p := res
	var tmp *ListNode
	for head != nil && center != nil {
		if head.Val < center.Val {
			tmp = head
			head = head.Next
		} else {
			tmp = center
			center = center.Next
		}
		tmp.Next = nil
		p.Next = tmp
		p = p.Next
	}
	if head != nil {
		p.Next = head
	}
	if center != nil {
		p.Next = center
	}
	return res.Next
}

func main() {
	nums := []int{4, 2, 1, 3, 0}
	head := CreateList(nums)
	head = sortList(head)
	PrintList(head)
}
