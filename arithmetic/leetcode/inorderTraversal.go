package main

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func main() {

}

func inorderTraversal(root *TreeNode) []int {
	rue := make([]int, 0)
	pNode := root
	tmp := make([]*TreeNode, 0)
	for pNode != nil || len(tmp) != 0 {
		if pNode != nil {
			tmp = append(tmp, pNode)
			pNode = pNode.Left
		} else {
			node := pNode
			if len(tmp) > 0 {
				node = tmp[len(tmp)-1]
				tmp = tmp[:len(tmp)-1]
			}
			rue = append(rue, node.Val)
			pNode = node.Right
		}
	}
	return rue
}
