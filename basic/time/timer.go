package main

import (
	"fmt"
	"sync"
	"time"
)

func main() {
	d := time.Second * 10
	t := time.NewTimer(d)
	wg := sync.WaitGroup{}
	t1 := time.Now()
	wg.Add(1)
	go func() {
		select {
		case <-t.C:
			fmt.Println("hello")
			wg.Done()
		}
	}()
	wg.Wait()
	fmt.Println(time.Since(t1))
}
