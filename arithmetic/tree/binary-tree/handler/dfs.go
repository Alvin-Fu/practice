package handler

import (
    . "practice/arithmetic/tree/binary-tree/model"
)
//深度遍历
func DFS(root *TreeNode, result *[]int){
    if root == nil {
        return
    }
    *result = append(*result, root.Val)
    DFS(root.Left, result)
    DFS(root.Right, result)
}

func DivideAndConquer(root *TreeNode)[]int{
    if root == nil {
        return nil
    }
    result := make([]int, 0)
    left := DivideAndConquer(root.Left)
    right := DivideAndConquer(root.Right)
    result = append(result, root.Val)
    result = append(result, left...)
    result = append(result, right...)
    return result
}
