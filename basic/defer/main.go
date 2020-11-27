package main

import "fmt"

func main(){
    fmt.Println(f1(),f2(), f3())
}

func f1() (result int) {
    defer func() {
        result++
    }()
    return 0
}

func f2() (r int) {
    t := 5
    defer func() {
        t = t + 5
        fmt.Println(t)
    }()
    return t
}

func f3() (r int) {
    defer func(r int) {
        r = r + 5
    }(r)
    return 1
}
