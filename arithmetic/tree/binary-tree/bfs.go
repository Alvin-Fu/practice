package main

import (
    . "practice/arithmetic/tree/binary-tree/model"
    	"practice/arithmetic/tree/binary-tree/handler"
    "fmt"
)
// 广度优先遍历，需要借助队列进行
func main() {
    handler.Bfs(Root)
    result := handler.LevelOrder(Root)
    fmt.Println(result)
}

