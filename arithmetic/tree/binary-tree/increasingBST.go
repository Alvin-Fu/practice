package main

import "practice/arithmetic/tree/binary-tree/model"

func increasingBST(head *model.TreeNode) *model.TreeNode {
	inorder(head)
	return resNode.Right
}

var resNode = &model.TreeNode{}

func inorder(node *model.TreeNode) {
	if node == nil {
		return
	}
	inorder(node.Left)
	node.Right = node
	node.Left = nil
	resNode = node
	inorder(node.Right)
}
