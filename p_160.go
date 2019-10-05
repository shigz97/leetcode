package main

func getIntersectionNode(headA, headB *ListNode) *ListNode {
	dic := make(map[*ListNode]struct{})
	var mem struct{}
	for headA != nil {
		dic[headA] = mem
		headA = headA.Next
	}
	for headB != nil {
		if _, ok := dic[headB]; ok {
			return headB
		}
		headB = headB.Next
	}
	return nil
}
