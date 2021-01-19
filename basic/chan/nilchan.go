package main

import (
	"fmt"
	"time"
)

type NilChan struct {
	c chan int
}

func main() {
	nilTest2()
}

func nilTest1() {
	c := NilChan{}
	go func() {
		defer fmt.Println("kk")
		c.c <- 1
		fmt.Println("hello")
	}()
	time.Sleep(1 * time.Second)
	c.c = make(chan int)
	go func() {
		c.c <- 1
	}()
	go func() {
		fmt.Println(<-c.c)
		fmt.Println(<-c.c)
	}()
	time.Sleep(5 * time.Second)
	close(c.c)
}

func nilTest2() {
	ch1 := make(chan int)
	var ch2 chan int
	go func() {
		for {
			select {
			case i := <-ch1:
				fmt.Println("1", i)
				ch2 = ch1
			case j := <-ch2:
				fmt.Println(j)
			}
		}
	}()
	go func() {
		ch2 <- 1
	}()
	time.Sleep(1 * time.Second)
	ch1 <- 2
	time.Sleep(3 * time.Second)
}
