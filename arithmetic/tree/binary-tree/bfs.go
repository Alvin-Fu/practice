package main

import (
    . "practice/arithmetic/tree/binary-tree/model"
    "fmt"
)
// 广度优先遍历，需要借助队列进行
func main() {
    bfs(Root)
}

func bfs(root *TreeNode) {
	if root == nil {
		return
	}
	queue := make([]*TreeNode, 0)
	result := make([]int, 0)
	queue = append(queue, root)
	for  len(queue) != 0 {
	    temp := queue[0]
        queue = queue[1:]
		result = append(result, temp.Val)
		if temp.Left != nil {
			queue = append(queue, temp.Left)
		}
        if temp.Right != nil {
            queue = append(queue, temp.Right)
        }
	}
    for _, r := range result{
        fmt.Printf("val: %d\t", r)
    }
    fmt.Printf("\n")
}
