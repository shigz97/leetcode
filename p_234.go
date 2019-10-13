package main

import (
	"fmt"
	. "leetcode/linklist"
)

func isPalindrome(head *ListNode) bool {
	if head == nil {
		return true
	}
	slow, fast := head, head.Next
	if fast == nil {
		return true
	}
	stack,top := []*ListNode,-1
	
}

func main() {
	nums := []int{1, 2, 2, 1}
	head := CreateList(nums)
	fmt.Println(isPalindrome(head))
}
