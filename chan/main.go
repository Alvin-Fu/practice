package main

import (
	"fmt"
	"sync"
	"time"
)

func main() {
	var ch = make(chan struct{})
	wg := sync.WaitGroup{}
	wg.Add(1)
	go func() {
		time.Sleep(5 * time.Second)
		close(ch)
		wg.Done()
	}()
	for {
		select {
		case <-ch:
			goto exit
		}
	}
exit:
	fmt.Println("bye")
}
