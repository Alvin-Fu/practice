package model

type TreeNode struct {
    Val   int
    Left  *TreeNode
    Right *TreeNode
}
var Root *TreeNode
func init(){
    left7 := &TreeNode{
        Val: 12,
    }
    right6 := &TreeNode{
        Val: 10,
    }
    left6 := &TreeNode{
        Val: 11,
        Left: left7,
    }
    right5 := &TreeNode{
        Val: 9,
        Right: right6,
        Left: left6,
    }
    right4 := &TreeNode{
        Val: 8,
    }
    left3:= &TreeNode{
        Val: 6,
        Right: right4,
    }
    right3 := &TreeNode{
        Val: 7,
    }
    left2 := &TreeNode{
        Val: 4,
    }
    right2 := &TreeNode{
        Val: 5,
        Right: right5,
    }
    left1 := &TreeNode{
        Val: 2,
        Left: left2,
        Right: right2,
    }
    right1 := &TreeNode{
        Val: 3,
        Left: left3,
        Right: right3,
    }
    Root = &TreeNode{
        Val: 1,
        Left: left1,
        Right: right1,
    }
}
