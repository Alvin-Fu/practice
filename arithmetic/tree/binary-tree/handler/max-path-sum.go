package handler

import (
    "practice/arithmetic/tree/binary-tree/model"
    )


// 获取二叉树的最大路径和
func MaxPathSum(root *model.TreeNode)int{
    if root == nil {
        return 0
    }
    result := maxPathSum(root)
    return result.MaxPath
}
// 用于保存单边最大的值和双边最大值
type Result struct {
    SinglePath int
    MaxPath int
}

func maxPathSum(root *model.TreeNode) *Result{
    if root == nil {
        return  &Result{
            SinglePath: 0,
            MaxPath: -(1<<31),
        }
    }
    left := maxPathSum(root.Left)
    right := maxPathSum(root.Right)
    result := &Result{}
    if left.SinglePath > right.SinglePath {
        result.SinglePath = max(left.SinglePath + root.Val, 0)
    } else {
        result.SinglePath = max(right.SinglePath + root.Val, 0)
    }
    maxPath := max(left.MaxPath, right.MaxPath)
    result.MaxPath = max(maxPath, left.SinglePath + right.SinglePath + root.Val)
    return result
}

func max(a, b int)int{
    if a > b {
        return  a
    }
    return  b
}