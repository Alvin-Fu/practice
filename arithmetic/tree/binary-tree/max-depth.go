package main

import (
    "practice/arithmetic/tree/binary-tree/model"
    "practice/arithmetic/tree/binary-tree/handler"
    "fmt"
)

func main(){
    depth := handler.MaxDepth(model.Root)
    fmt.Println(depth)
}


