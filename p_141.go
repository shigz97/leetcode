package main

import (
	"fmt"
	. "leetcode/linklist"
)

func hasCycle(head *ListNode) bool {
	if head == nil || head.Next == nil {
		return false
	}
	slow, fast := head, head.Next
	for slow != fast {
		if fast == nil || fast.Next == nil {
			return false
		}
		slow = slow.Next
		fast = fast.Next.Next
	}
	return true
}

func main() {
	nums := []int{3, 2, 0, -4}
	head := CreateCycleList(nums, 1)
	fmt.Println(hasCycle(head))
}
