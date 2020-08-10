package main

import (
    ."practice/arithmetic/tree/binary-tree/model"
    "fmt"
)

func main(){
    recursionI(Root)
    fmt.Printf("\n")
    cyclicI(Root)
}

func recursionI(root *TreeNode){
    if root == nil {
        return
    }
    recursionI(root.Left)
    fmt.Printf("val: %d\t", root.Val)
    recursionI(root.Right)
}

func cyclicI(root *TreeNode){
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
