package main

import (
	"fmt"
	"time"
)

func main() {
	test()

}

func test() {
	ch1 := make(chan int)
	go func() {
		fmt.Println(<-ch1)
	}()
	ch1 <- 5
	time.Sleep(1 * time.Second)
}

func test2() {
	c := make(chan int)
	go func() {
		select {
		case c <- 1:

		}
	}()
	time.Sleep(1 * time.Second)
	go func() {
		select {
		case c <- 1:

		}
	}()
	time.Sleep(3 * time.Second)
	fmt.Println(<-c)
}
