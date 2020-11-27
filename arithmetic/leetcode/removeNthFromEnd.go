package main

import (
	. "practice/arithmetic/leetcode/model"
)

func removeNthFromEnd(head *ListNode, n int) *ListNode {
	tmp := make(map[int]*ListNode)
	node := head
	index := 0
	for node.Next != nil {
		tmp[index] = node
		index++
		node = node.Next
	}
	tmp[index] = node
	i := len(tmp) - n
	tmp[i-1].Next = tmp[i+1]
	tmp[1].Next = nil
	return head
}
