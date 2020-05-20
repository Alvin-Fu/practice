package main

type ListNode struct {
	Val int
	Next *ListNode
}

func AddTwoNumbers(l1, l2 *ListNode)*ListNode{
	var rue = new(ListNode)
	if l1 == nil {
		return l2
	}
	if l2 == nil {
		return l1
	}
	tmp1, tmp2 := l1, l2
	tmp := rue
	i := 0
	for tmp1 != nil || tmp2 != nil{
		t1, t2 := 0, 0
		if tmp1 != nil{
			t1 = tmp1.Val
		}
		if tmp2 != nil  {
			t2 = tmp2.Val
		}
		sum := t1 + t2 + i
		tmp.Val = sum %10
		i = sum / 10
		if tmp1 != nil {
			tmp1 = tmp1.Next
		}
		if tmp2 != nil {
			tmp2 = tmp2.Next
		}
			tmp.Next = &ListNode{}
			tmp = tmp.Next
	}
	if i != 0{
		if tmp == nil {
			tmp.Next = &ListNode{}
			tmp = tmp.Next
		}
		tmp.Val = 1
	} else {
		tmp.Next = nil
	}
	return rue
}

