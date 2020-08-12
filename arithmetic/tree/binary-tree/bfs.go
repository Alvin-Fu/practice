package main

import (
    . "practice/arithmetic/tree/binary-tree/model"
    	"practice/arithmetic/tree/binary-tree/handler"
)
// 广度优先遍历，需要借助队列进行
func main() {
    handler.Bfs(Root)
}

