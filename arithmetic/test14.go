package main

import "golang.org/x/tools/go/ssa/interp/testdata/src/fmt"

type TreeNode struct {
	Val   int
	left  *TreeNode
	right *TreeNode
}

// 查看前十名
func main() {
	tmp := []int{}
}

func findTen(root *TreeNode) {
	que := make([]*TreeNode, 0)
	rue := make([]int, 0)
	que = append(que, root)
	for len(que) != 0 {
		tmp := que[0]
		que = que[1:]
		rue = append(rue, tmp.Val)
		if tmp.left != nil {
			que = append(que, tmp.left)
		}
		if tmp.right != nil {
			que = append(que, tmp.right)
		}
	}
	fmt.Println(rue)
}
