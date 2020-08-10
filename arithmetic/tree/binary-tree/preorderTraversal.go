package main

import (
    "fmt"
    ."practice/arithmetic/tree/binary-tree/model"
    )


// 前序遍历
func main() {
    cyclicP(Root)
    fmt.Printf("\n")
    recursionP(Root)
    fmt.Printf("\n")
}

func recursionP(root *TreeNode){
    if root == nil {
        return
    }
    fmt.Printf("val: %d\t", root.Val)
    recursionP(root.Left)
    recursionP(root.Right)
}

func cyclicP(root *TreeNode){
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

func cyclicP2(root *TreeNode){
    if root == nil {
        return
    }
    result := make([]int, 0)
    stack := make([]*TreeNode, 0)
    temp := root
}
