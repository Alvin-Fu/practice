package main

import (
	"fmt"
	"time"
)

func main() {
	var ch = make(chan struct{})
	//var c = make(chan struct{})
	go func() {
		time.Sleep(1 * time.Second)
		ch <- struct{}{}
	}()
	c := fanIn(ch)
	<-c
	fmt.Println("bye")
//	wg := sync.WaitGroup{}
//	wg.Add(1)
//	go func() {
//		time.Sleep(5 * time.Second)
//		close(ch)
//		wg.Done()
//	}()
//	for {
//		select {
//		case <-ch:
//			goto exit
//		}
//	}
//exit:
//	fmt.Println("bye")
}

func fanIn(ch chan struct{})chan struct{}{
	c := make(chan struct{})
	go func() {
		c <- <- ch
	}()
	return c
}