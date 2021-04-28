package main

import (
	"runtime"
	"sync"
	"fmt"
)

func main(){
	fmt.Println(runtime.NumCPU())
	runtime.GOMAXPROCS(1)
	wg := sync.WaitGroup{}
	wg.Add(12)
	for i := 0; i < 6; i++{
		go func() {
			fmt.Println("T1: ", i)
			wg.Done()
		}()
	}
	for i := 0; i < 6; i++{
		go func(i int) {
			fmt.Println("T2: ", i)
			wg.Done()
		}(i)
	}
	wg.Wait()
}
