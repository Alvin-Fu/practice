package handler

import (
    "practice/arithmetic/tree/binary-tree/model"
    )

func IsBalancedBTree(root *model.TreeNode)bool{
   if MaxDepthBT(root) == -1{
        return false
    }
    return true
}
