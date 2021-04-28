package main

import (
	"fmt"
	"practice/arithmetic/tree/binary-tree/model"
)

func main() {
	preorderTraversal(model.Root)
	preorderTraversal2(model.Root)
}

func preorderTraversal(root *model.TreeNode) {
	if root == nil {
		return
	}
	stack := make([]*model.TreeNode, 0)
	rue := make([]int, 0)
	tmp := root
	for {
		if tmp == nil {
			if len(stack) > 0 {
				tmp = stack[len(stack)-1].Right
				stack = stack[:len(stack)-1]
				continue
			} else {
				break
			}
		}
		rue = append(rue, tmp.Val)
		if tmp.Right != nil {
			stack = append(stack, tmp)
		}
		tmp = tmp.Left
	}
	fmt.Println(rue)
}

func preorderTraversal2(root *model.TreeNode) {
	if root == nil {
		return
	}
	stack := make([]*model.TreeNode, 0)
	rue := make([]int, 0)
	tmp := root
	for {
		if tmp == nil {
			if len(stack) > 0 {
				// 左节点已经遍历了，将节点出栈
				tmp = stack[len(stack)-1].Right
				stack = stack[:len(stack)-1]
				continue
			} else {
				break
			}
		}
		rue = append(rue, tmp.Val)
		// 右节点存在的时候将tmp压栈
		if tmp.Right != nil {
			stack = append(stack, tmp)
		}
		tmp = tmp.Left
	}
	fmt.Println(rue)
}
