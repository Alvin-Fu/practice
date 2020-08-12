package main

import (
    "practice/arithmetic/tree/binary-tree/handler"
    "practice/arithmetic/tree/binary-tree/model"
    "fmt"
)

func main(){
    sum := handler.MaxPathSum(model.Root)
    fmt.Println(sum)
}
