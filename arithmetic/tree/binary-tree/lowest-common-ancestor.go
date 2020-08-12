package main

import (
    "practice/arithmetic/tree/binary-tree/handler"
    "practice/arithmetic/tree/binary-tree/model"
    "fmt"
)

func main(){
    right4 := &model.TreeNode{
        Val: 4,
    }
    left7 := &model.TreeNode{
        Val: 6,
    }
    node := handler.LowestCommonAncestor(model.Root, right4, left7)
    fmt.Println(node)
}
