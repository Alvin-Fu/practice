package main

import (
	. "practice/arithmetic/tree/binary-tree/model"
    "fmt"
    "practice/arithmetic/tree/binary-tree/handler"
)

// 后续遍历
func main() {
    handler.RecursionPo(Root)
    fmt.Printf("\n")
    handler.CyclicPo(Root)
}

