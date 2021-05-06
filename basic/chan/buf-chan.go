package main

import (
	"fmt"
	"time"
)

func main() {
	ch := make(chan int64, 2)
	ch <- 1
	ch <- 2
	go func() {
		ch <- 3
		fmt.Println("hello")
	}()
	time.Sleep(time.Second * 1)
	for i := 0; i < 3; i++ {
		fmt.Println(<-ch)
	}
}
