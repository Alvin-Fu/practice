package main

import (
    "fmt"
    ."practice/arithmetic/tree/binary-tree/model"
    "practice/arithmetic/tree/binary-tree/handler"
)


// 前序遍历
func main() {
    handler.CyclicP(Root)
    fmt.Printf("\n")
    handler.RecursionP(Root)
    fmt.Printf("\n")
}

