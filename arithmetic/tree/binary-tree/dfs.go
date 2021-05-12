package main

import (
    . "practice/arithmetic/tree/binary-tree/model"
    "fmt"
    "practice/arithmetic/tree/binary-tree/handler"
)

func main(){
    result := make([]int, 0)
    handler.DFS(Root, &result)
    for _, r := range result{
        fmt.Printf("val: %d\t", r)
    }
    fmt.Printf("\n")
    result = handler.DivideAndConquer(Root)
    for _, r := range result{
        fmt.Printf("val: %d\t", r)
    }
    fmt.Printf("\n")
}

