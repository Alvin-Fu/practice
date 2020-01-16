package main

import (
	"time"
	"fmt"
)

func main(){
	ch := make(chan int)
	out := make(chan int)
	go producer(ch, 100 * time.Millisecond)
	go producer(ch, 200 *time.Millisecond)
	go reader(out)
	for {
		x, ok := <-ch
		if !ok {
			break
		}
		out <- x
	}

}

func producer(ch chan int, d time.Duration){
	var i int
	for {
		ch <- i
		i ++

		if i == 10{
			ch <- i
			break
		}
		time.Sleep(d)
	}
}

func reader(out chan int){
	for {
		x, ok := <-out
		if !ok {
			break
		}
		fmt.Println(x)
	}
}
