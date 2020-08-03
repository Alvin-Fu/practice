package main

import (
    "os"
    "runtime/trace"
    "fmt"
)

func main(){
    file, err := os.Create("trace.out")
    if err != nil {
        panic(err)
    }
    defer file.Close()
    err = trace.Start(file)
    if err != nil {
        panic(err)
    }
    defer trace.Stop()
    fmt.Println("hello trace")
}