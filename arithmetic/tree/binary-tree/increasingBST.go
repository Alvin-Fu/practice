package main

import "practice/arithmetic/tree/binary-tree/model"

var (
	t       *model.TreeNode
	resNode *model.TreeNode
)

func increasingBST(root *model.TreeNode) *model.TreeNode {
	t = new(model.TreeNode)
	resNode = t
	if root == nil {
		return root
	}
	inorder(root)
	return t.Right
}

func inorder(node *model.TreeNode) {
	if node == nil {
		return
	}
	inorder(node.Left)
	resNode.Right = node
	node.Left = nil
	resNode = node
	inorder(node.Right)
}
