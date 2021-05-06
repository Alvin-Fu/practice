package main

import (
	"fmt"
	"practice/arithmetic/tree/binary-tree/model"
)

func main() {
	postorderTraversal(model.Root)
	postorderTraversal2(model.Root)
}

func postorderTraversal(root *model.TreeNode) {
	if root == nil {
		return
	}
	stack := make([]*model.TreeNode, 0)
	rue := make([]int, 0)
	tmp := root
	lastNode := new(model.TreeNode)
	for tmp != nil || len(stack) != 0 {
		for tmp != nil {
			stack = append(stack, tmp)
			tmp = tmp.Left
		}
		tmp = stack[len(stack)-1]
		if tmp.Right == nil || tmp.Right == lastNode {
			rue = append(rue, tmp.Val)
			stack = stack[:len(stack)-1]
			lastNode = tmp
			tmp = nil
		} else {
			tmp = tmp.Right
		}
	}
	fmt.Println(rue)
}

func postorderTraversal2(root *model.TreeNode) {
	if root == nil {
		return
	}
	stack := make([]*model.TreeNode, 0)
	rue := make([]int, 0)
	tmp := root
	lastNode := new(model.TreeNode)
	for tmp != nil || len(stack) != 0 {
		for tmp != nil {
			stack = append(stack, tmp)
			tmp = tmp.Left
		}
		tmp = stack[len(stack)-1]
		if tmp.Right == nil || tmp.Right == lastNode {
			rue = append(rue, tmp.Val)
			stack = stack[:len(stack)-1]
			lastNode = tmp
			tmp = nil
		} else {
			tmp = tmp.Right
		}
	}
	fmt.Println(rue)
}
