package main

import (
	"sync"
	"time"
	"fmt"
	"github.com/petermattis/goid"
)

func main(){
	wg := sync.WaitGroup{}
	wg.Add(36)
	go pool(&wg, 36, 50)
	wg.Wait()
}

func worker(taskCh chan int, wg *sync.WaitGroup){
	defer wg.Done()
	for {
		task, ok := <- taskCh
		if !ok {
			return
		}
		d := time.Duration(task) * time.Millisecond
		time.Sleep(d)
		fmt.Printf("gid: %d, task: %d\n", goid.Get(), task)
	}

}

func pool(wg *sync.WaitGroup, workers, tasks int){
	taskCh := make(chan int)
	for i := 0; i < workers; i ++{
		go worker(taskCh, wg)
	}
	for i := 0; i < tasks; i++{
		taskCh <- i
	}
	close(taskCh)
}
