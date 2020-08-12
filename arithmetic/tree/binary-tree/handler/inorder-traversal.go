package handler


import (
    . "practice/arithmetic/tree/binary-tree/model"
    "fmt"
)

// 中序遍历
func RecursionI(root *TreeNode){
    if root == nil {
        return
    }
    RecursionI(root.Left)
    fmt.Printf("val: %d\t", root.Val)
    RecursionI(root.Right)
}

func CyclicI(root *TreeNode){
    if root == nil{
        return
    }
    stack := make([]*TreeNode, 0)
    result := make([]int, 0)
    temp := root
    for temp != nil || len(stack) != 0 {
        for temp != nil {
            stack = append(stack, temp)
            temp = temp.Left
        }
        temp = stack[len(stack) - 1]
        stack = stack[:len(stack) - 1]
        result = append(result, temp.Val)
        temp = temp.Right
    }

    for _, r := range result{
        fmt.Printf("val: %d\t", r)
    }
    fmt.Printf("\n")
}
