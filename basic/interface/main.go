package main

import (
    "fmt"
    )
//内部类型方法集提升的规则
type worker interface {
    work()
}

type person struct {
    name string
    worker
}

func main(){
    var w worker = person{}
    fmt.Println(w)
}
