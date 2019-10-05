package main

import "fmt"

type ListNode struct {
	key  int
	Val  int
	Pre  *ListNode
	Next *ListNode
}

type LRUCache struct {
	capacity int
	dic      map[int]*ListNode
	head     *ListNode
	tail     *ListNode
}

func removeNode(node *ListNode) {
	p1, p2 := node.Pre, node.Next
	p1.Next, p2.Pre = p2, p1
}

func addNode(node *ListNode, head *ListNode) {
	node.Pre, node.Next = head, head.Next
	head.Next.Pre, head.Next = node, node
}

func moveToHead(node *ListNode, head *ListNode) {
	removeNode(node)
	addNode(node, head)
}

func popTail(node *ListNode) *ListNode {
	p := node.Pre
	p.Pre.Next, node.Pre = node, p.Pre
	return p
}

func Constructor(capacity int) LRUCache {
	//return LRUCache{capacity, make(map[int]int, capacity)}
	head, tail := &ListNode{0, 0, nil, nil}, &ListNode{0, 0, nil, nil}
	head.Next, tail.Pre = tail, head
	return LRUCache{capacity, make(map[int]*ListNode, capacity), head, tail}
}

func (this *LRUCache) Get(key int) int {
	res, ok := this.dic[key]
	if !ok {
		return -1
	} else {
		moveToHead(res, this.head)
		return res.Val
	}
}

func (this *LRUCache) Put(key int, value int) {
	res, ok := this.dic[key]
	if ok {
		res.Val = value
		moveToHead(res, this.head)
	} else {
		if len(this.dic) == this.capacity {
			p := popTail(this.tail)
			delete(this.dic, p.key)
		}
		res := &ListNode{key, value, nil, nil}
		this.dic[key] = res
		addNode(res, this.head)
	}
}

func main() {
	obj := Constructor(2)
	fmt.Println(obj.Get(1))
	obj.Put(1, 2)
	fmt.Println(obj.Get(1))
	obj.Put(1, 3)
	fmt.Println(obj.Get(1))
	obj.Put(2, 4)
	obj.Put(3, 5)
	fmt.Println(obj.dic)
}
