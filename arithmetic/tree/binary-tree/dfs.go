package main

import (
    . "practice/arithmetic/tree/binary-tree/model"
    "fmt"
)

func main(){
    result := make([]int, 0)
    dfs(Root, &result)
    for _, r := range result{
        fmt.Printf("val: %d\t", r)
    }
    fmt.Printf("\n")
    result = divideAndConquer(Root)
    for _, r := range result{
        fmt.Printf("val: %d\t", r)
    }
    fmt.Printf("\n")
}

func dfs(root *TreeNode, result *[]int){
    if root == nil {
        return
    }
    *result = append(*result, root.Val)
    dfs(root.Left, result)
    dfs(root.Right, result)
}

func divideAndConquer(root *TreeNode)[]int{
    if root == nil {
        return nil
    }
    result := make([]int, 0)
    left := divideAndConquer(root.Left)
    right := divideAndConquer(root.Right)
    result = append(result, root.Val)
    result = append(result, left...)
    result = append(result, right...)
    return result
}
