package main

import (
	"fmt"
)

/*
给定一个单向链表，节点数据结构如下，除了包含一个next指针，还有一个rand指针，
随机指向链表中的某一元素，或者为空指针。实现一个函数，深复制这个链表。
*/

type Node struct {
	Val  int
	next *Node
	rand *Node
}

func main() {
	//left7 := &Node{
	//	Val: 12,
	//}
	//right6 := &Node{
	//	Val: 10,
	//}
	//left6 := &Node{
	//	Val: 11,
	//	Left: left7,
	//}
	//right5 := &Node{
	//	Val: 9,
	//	Right: right6,
	//	Left: left6,
	//}
	//right4 := &Node{
	//	Val: 8,
	//}
	//left3:= &Node{
	//	Val: 6,
	//	Right: right4,
	//}
	//right3 := &Node{
	//	Val: 7,
	//}
	left2 := &Node{
		Val: 4,
	}
	right2 := &Node{
		Val:  5,
		next: left2,
	}
	left1 := &Node{
		Val:  2,
		rand: left2,
		next: right2,
	}
	right1 := &Node{
		Val:  3,
		next: left1,
		rand: right2,
	}
	Root := &Node{
		Val:  -10,
		next: right1,
		rand: left1,
	}
	h := Test(Root)
	for h != nil {

		if h.rand != nil {
			fmt.Printf("rand: %v ", *h.rand)
		}
		fmt.Println(*h)
		h = h.next
	}
}

func Test(head *Node) *Node {
	if head == nil {
		return head
	}
	newHead := &Node{
		Val: head.Val,
	}
	node2 := newHead
	tmp := make(map[*Node]*Node)
	if head.rand != nil {
		tmp[head.rand] = newHead
	}
	node := head.next
	for node != nil {
		node2.next = &Node{
			Val: node.Val,
		}
		if node.rand != nil {
			tmp[node.rand] = node2.next
		}
		if v, ok := tmp[node]; ok {
			v.rand = node2.next
		}
		node = node.next
		node2 = node2.next
	}
	return newHead
}
