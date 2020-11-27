package main

import (
        "fmt"
    . "practice/arithmetic/leetcode/model"
)



func main(){
    node3 := &ListNode{
        Val: 4,
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
    tmp := swapPairs(head)

    for tmp != nil {
        fmt.Println(tmp.Val)
        tmp = tmp.Next
    }
}

func swapPairs(head *ListNode)*ListNode{
    if head == nil || head.Next == nil {
        return head
    }
    tmp := head
    temp := head
    for i := 0; tmp != nil && tmp.Next != nil; i ++ {
        node1, node2 := tmp, tmp.Next
        if i != 0 {
            temp.Next = node2
            temp = node1
        }
        tmp = node2.Next
        node1.Next = node2.Next
        node2.Next = node1

    }

    return head.Next
}
