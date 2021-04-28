package main

import (
	"fmt"
	"practice/arithmetic/tree/binary-tree/model"
)

func main() {
	inorderTraversal(model.Root)
	inorderTraversal2(model.Root)
}

func inorderTraversal(root *model.TreeNode) {
	if root == nil {
		return
	}
	stack := make([]*model.TreeNode, 0)
	rue := make([]int, 0)
	//stack = append(stack, root)
	tmp := root
	for tmp != nil || len(stack) != 0 {
		for tmp != nil {
			stack = append(stack, tmp)
			tmp = tmp.Left
		}
		tmp = stack[len(stack)-1]
		stack = stack[:len(stack)-1]
		rue = append(rue, tmp.Val)
		tmp = tmp.Right
	}
	fmt.Println(rue)
}

func inorderTraversal2(root *model.TreeNode) {
	if root == nil {
		return
	}
	stack := make([]*model.TreeNode, 0)
	rue := make([]int, 0)
	tmp := root
	for tmp != nil || len(stack) > 0 {
		for tmp != nil {
			stack = append(stack, tmp)
			tmp = tmp.Left
		}
		tmp = stack[len(stack)-1]
		rue = append(rue, tmp.Val)
		stack = stack[:len(stack)-1]
		tmp = tmp.Right
	}
	fmt.Println(rue)
}
