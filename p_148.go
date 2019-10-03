package main

import (
	. "leetcode/linklist"
)

func sortList(head *ListNode) *ListNode {

}

func main() {
	nums := []int{4, 2, 1, 3}
	head := CreateList(nums)
	head = sortList(head)
	PrintList(head)
}
