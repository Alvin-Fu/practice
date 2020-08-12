package handler

import (
    "practice/arithmetic/tree/binary-tree/model"
    "fmt"
)

func LowestCommonAncestor(root, p, q *model.TreeNode) *model.TreeNode{
    if root == nil || p == nil || q == nil  {
        return nil
    }
    if root.Val == p.Val || root.Val == q.Val{
        return root
    }
    left := LowestCommonAncestor(root.Left, p, q)
    right := LowestCommonAncestor(root.Right, p, q)
    if left != nil && right != nil  {
        return root
    }
    if left != nil {
        return left
    }
    if right != nil {
        return right
    }
    return nil
}

type CommonResult struct {
    common *model.TreeNode
    isP, isQ bool
}

func lowest(root, p, q *model.TreeNode)*CommonResult{
    if root == nil{
        return nil
    }

    lowest(root.Left, p, q)
    lowest(root.Right, p, q)
    if root.Val == p.Val {

        fmt.Println(root.Val)
    }
    if root.Val == q.Val {
        fmt.Println(root.Val)
    }
    return nil
}
