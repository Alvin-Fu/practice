package main

import (
	"time"
	"fmt"
	"github.com/petermattis/goid"
)

func main() {
	var Ball int
	table := make(chan int)
	for i := 0; i < 10; i++{
		go player(table)
	}
	table <- Ball
	time.Sleep(1 * time.Second)
	<-table
}

func player(table chan int) {
	for {
		ball := <-table
		ball++
		fmt.Printf("gid: %d, ball: %d\n",goid.Get(),ball)
		time.Sleep(100 * time.Millisecond)
		table <- ball
	}
}
