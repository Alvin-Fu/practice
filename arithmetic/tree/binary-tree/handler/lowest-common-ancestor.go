package handler

import (
    "practice/arithmetic/tree/binary-tree/model"
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

func Lowest(root, p, q *model.TreeNode)*CommonResult{
    if root == nil{
        return nil
    }
    result := &CommonResult{}
    if root.Val == p.Val {
        result.common = root
        result.isP = true
    }
    if root.Val == q.Val {
        result.common = root
        result.isQ = true
    }
    left := Lowest(root.Left, p, q)
    right := Lowest(root.Right, p, q)
    if left.isP && left.isQ {
        return  left
    }
    if right.isP &&right.isQ{
        return right
    }
    return nil
}
