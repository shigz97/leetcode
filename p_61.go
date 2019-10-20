package main

import (
	. "github.com/shigz97/leetcode/linklist"
)

func rotateRight(head *ListNode, k int) *ListNode {
	if head == nil {
		return nil
	}
	fast, slow := head, head
	for k > 0 {
		fast = fast.Next
		if fast == nil {
			fast = head
		}
		k--
	}
	//fmt.Println(fast, slow)
	for fast.Next != nil {
		slow, fast = slow.Next, fast.Next
	}
	if slow.Next == nil {
		return head
	}
	newhead := slow.Next
	fast.Next = head
	slow.Next = nil
	return newhead
	//return nil
}

func main() {
	nums := []int{1, 2, 3, 4, 5}
	head := CreateList(nums)
	head = rotateRight(head, 0)
	PrintList(head)
}
