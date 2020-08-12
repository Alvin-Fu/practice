package handler

import (
    . "practice/arithmetic/tree/binary-tree/model"
    "fmt"
)
// 前序遍历
func RecursionP(root *TreeNode){
    if root == nil {
        return
    }
    fmt.Printf("val: %d\t", root.Val)
    RecursionP(root.Left)
    RecursionP(root.Right)
}

func CyclicP(root *TreeNode){
    if root == nil {
        return
    }
    stack := make([]*TreeNode, 0)
    tmp := root
    for {
        if tmp == nil  {
            if len(stack) != 0 {
                tmp = stack[len(stack) - 1].Right
                stack = stack[:len(stack) - 1]
                continue
            } else {
                break
            }
        }
        fmt.Printf("val: %d\t", tmp.Val)
        if tmp.Right != nil {
            stack = append(stack, tmp)
        }
        tmp = tmp.Left
    }

}

