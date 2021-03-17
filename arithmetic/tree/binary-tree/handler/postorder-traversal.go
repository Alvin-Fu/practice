package handler

import (
	"fmt"
	. "practice/arithmetic/tree/binary-tree/model"
)

// 后续遍历
func RecursionPo(root *TreeNode) {
	if root == nil {
		return
	}
	RecursionPo(root.Left)
	RecursionPo(root.Right)
	fmt.Printf("val: %d\t", root.Val)
}

func CyclicPo(root *TreeNode) {
	if root == nil {
		return
	}
	stack := make([]*TreeNode, 0)
	result := make([]int, 0)
	temp := root
	lastNode := new(TreeNode)
	for temp != nil || len(stack) != 0 {
		for temp != nil {
			stack = append(stack, temp)
			temp = temp.Left
		}
		temp = stack[len(stack)-1]
		if temp.Right == nil || temp.Right == lastNode {
			stack = stack[:len(stack)-1]
			result = append(result, temp.Val)
			lastNode = temp
			//这块需要注意循环结束条件，需要将temp指向空，因为此时的temp已经遍历了
			temp = nil
		} else {
			temp = temp.Right
		}
	}
	for _, r := range result {
		fmt.Printf("val: %d\t", r)
	}
	fmt.Printf("\n")

}
