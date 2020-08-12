package main

import (
    ."practice/arithmetic/tree/binary-tree/model"
    "fmt"
    "practice/arithmetic/tree/binary-tree/handler"
)
func main(){
    handler.RecursionI(Root)
    fmt.Printf("\n")
    handler.CyclicI(Root)
}
