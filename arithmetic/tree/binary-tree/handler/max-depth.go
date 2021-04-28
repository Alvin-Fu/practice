package handler

import (
    "practice/arithmetic/tree/binary-tree/model"
    )

func MaxDepthBT(root *model.TreeNode)int {
    if root == nil {
        return 0
    }
    left := MaxDepthBT(root.Left)
    right := MaxDepthBT(root.Right)
    if left == -1 || right == -1 || left - right > 1 || right -left > 1{
        return -1
    }
    if left > right{
        return left + 1
    }
    return right + 1
}

func MaxDepth(root *model.TreeNode)int{
    if root == nil {
        return 0
    }
    left := MaxDepth(root.Left)
    right := MaxDepth(root.Right)
    if left > right{
        return left + 1
    }
    return right + 1
}


func DivideAndConquerDepth(root *model.TreeNode)int{
    if root == nil {
        return 0
    }
    depthLeft := DivideAndConquerDepth(root.Left)
    depthRight := DivideAndConquerDepth(root.Right)
    if depthLeft > depthRight{
       return depthLeft +1
    }
    return   depthRight + 1


}