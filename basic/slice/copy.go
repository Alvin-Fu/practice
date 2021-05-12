package main

import "fmt"

func main(){
    temp1 := make([]int, 10)
    temp2 := []int{1, 2,3, 4}
    copy(temp1[1:], temp2)
    fmt.Println(temp1)
}
