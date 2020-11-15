package main

import "fmt"

type ListNode struct {
	Val  int
	Next *ListNode
}

func main() {
	node4 := &ListNode{
		Val: 5,
	}
	node3 := &ListNode{
		Val:  4,
		Next: node4,
	}
	node2 := &ListNode{
		Val:  0,
		Next: node3,
	}
	node1 := &ListNode{
		Val:  2,
		Next: node2,
	}
	node4.Next = node3
	head := &ListNode{
		Val:  3,
		Next: node1,
	}
	detectCycle(head)
}

func detectCycle(head *ListNode) *ListNode {
	if head == nil || head.Next == nil {
		return nil
	}
	tmp1 := head.Next
	tmp2 := head.Next.Next
	index := 1
	for {
		if tmp1 == nil || tmp2 == nil || tmp2.Next == nil {
			return nil
		}
		if tmp1 == tmp2 {
			break
		}
		tmp1 = tmp1.Next
		tmp2 = tmp2.Next.Next
		index++
	}

	fmt.Println(index)
	return nil
}
