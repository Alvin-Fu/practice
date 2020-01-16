package main

import "sync"

const (
	Workers = 5
	SubWorker = 3
	Tasks = 20
	SubTasks = 10
)

func main(){
	wg := sync.WaitGroup{}
	wg.Add(Workers)
	tasks := make(chan int)
	for i := 0; i < Workers; i ++{
		go worker2(tasks, &wg)
	}
	for i := 0; i < Tasks; i++{
		tasks <- i
	}
	close(tasks)
	wg.Wait()
}

func worker2(tasks <- chan int, wg *sync.WaitGroup){
	defer wg.Done()
	for {
		task, ok := <-tasks
		if !ok {
			return
		}
		subTask := make(chan int)
		for i := 0; i < SubWorker; i ++{
			go subWorker(subTask)
		}
		for i := 0; i < SubTasks; i++{
			t := task * i
			subTask <- t
		}
		close(subTask)
	}
}
func subWorker(task chan int){

}
