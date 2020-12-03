package main

import "practice/arithmetic/leetcode/model"

/*
输入一个链表，输出该链表中倒数第k个节点。为了符合大多数人的习惯，本题从1开始计数，即链表的尾节点是倒数第1个节点。
例如，一个链表有6个节点，从头节点开始，它们的值依次是1、2、3、4、5、6。这个链表的倒数第3个节点是值为4的节点。

给定一个链表: 1->2->3->4->5, 和 k = 2.

返回链表 4->5.

*/

func main() {}

// 使用的快慢指针
func getKthFromEnd(head *model.ListNode, k int) *model.ListNode {
	if head == nil || head.Next == nil || k <= 0 {
		return head
	}
	temp1 := head
	temp2 := head
	for i := 0; i < k; i++ {
		temp2 = temp2.Next
	}
	for temp2 != nil {
		temp1 = temp1.Next
		temp2 = temp2.Next
	}
	return temp1
}
