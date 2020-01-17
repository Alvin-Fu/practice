package main

import (
	"fmt"
	)

func main(){
	ch := make(chan int)
	go generate(ch)
	for i := 0; i < 5; i ++{
		prime := <- ch
		fmt.Println(prime)
		ch1 := make(chan int)
		go filter(ch, ch1, prime)
		ch = ch1
	}
}

func generate(ch chan<- int){
	for i := 2; ; i++{
		//fmt.Printf("gid: %d, i: %d\n", goid.Get(), i)
		ch <- i
	}
}

func filter(in <-chan int, out chan<- int, prime int){
	for {
		i := <- in
		fmt.Printf("index: %d, prime: %d  \t",i, prime)
		if i % prime != 0 {
			fmt.Println()
			out <- i
		}
	}
}
