package handler

import (
    . "practice/arithmetic/tree/binary-tree/model"
    )

func Bfs(root *TreeNode) []int{
    if root == nil {
        return nil
    }
    queue := make([]*TreeNode, 0)
    result := make([]int, 0)
    queue = append(queue, root)
    for  len(queue) != 0 {
        temp := queue[0]
        queue = queue[1:]
        result = append(result, temp.Val)
        if temp.Left != nil {
            queue = append(queue, temp.Left)
        }
        if temp.Right != nil {
            queue = append(queue, temp.Right)
        }
    }
  return result
}

func LevelOrder( root *TreeNode)[][]int{
    result := make([][]int, 0)
    if root == nil {
        return result
    }
    queue := make([]*TreeNode, 0)
    queue = append(queue, root)
    for  len(queue) != 0 {
        list := make([]int, 0)
        l := len(queue)
        for i := 0; i < l; i++{
            temp := queue[0]
            queue = queue[1:]
            list = append(list, temp.Val)
            if temp.Left != nil {
                queue = append(queue, temp.Left)
            }
            if temp.Right != nil {
                queue = append(queue, temp.Right)
            }
        }
        result = append(result, list)
    }
    return result
}

