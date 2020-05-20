package main

import (
        "fmt"
)

func main()  {
    fmt.Println(reverse(-1534236469))

}

func reverse(x int) int {
    t := 0

    return t
}

func getNum(x int)int{
    var num int
    for i := 0; x / 10 > 0; i++{
        t := x % 10
        num = num * 10 + t
        x = x/10
    }
    return num
}
