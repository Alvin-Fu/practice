package main

<<<<<<< HEAD
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
=======
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
>>>>>>> 41655470a8bcee9960daa4f50479fc58afd85bf3
	}
	fmt.Println(rue)
}
