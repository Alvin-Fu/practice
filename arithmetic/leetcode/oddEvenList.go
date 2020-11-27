package main

import (
	. "practice/arithmetic/leetcode/model"
	"fmt"
	)

func main()  {
	node3 := &ListNode{
		Val: 4,
		Next: nil,
	}
	node2 := &ListNode{
		Val: 3,
		Next: node3,
	}
	node1 := &ListNode{
		Val: 2,
		Next: node2,
	}

	head := &ListNode{
		Val: 1,
		Next: node1,
	}
	//tmp := head
	//for tmp != nil {
	//	fmt.Println(*tmp)
	//	tmp = tmp.Next
	//}
	//return
	oddEvenList(head)
	//for head != nil {
	//	fmt.Println(*head)
	//	head = head.Next
	//}
}

func oddEvenList(head *ListNode) *ListNode {
	if head.Next == nil {
		return head
	}
	index := 0
	for head != nil {
		if index % 2 == 0{
			fmt.Println("偶数",*head)
		} else {
			fmt.Println(*head)
		}
		head = head.Next
		index ++
	}
	return head
}
