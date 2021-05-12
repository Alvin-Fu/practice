package main

import (
    "fmt"
    )

func main(){
    playing := []int{1, 2}
    ready := []int{3}
    tmp := make([]int, 0)
    tmp = append(tmp, playing...)
    tmp = append(tmp, ready...)
    for i, t := range tmp{
        fmt.Println(i, t)
    }
    fmt.Println("index: ", tmp[0])

}
