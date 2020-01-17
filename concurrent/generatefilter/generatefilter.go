package main

import (
	"fmt"
	"github.com/petermattis/goid"
)

func main(){
	ch := make(chan int)
	go generate(ch)
	for i := 0; i < 3; i ++{
		prime := <- ch
		fmt.Println("prime: ",prime)
		ch1 := make(chan int)
		go filter(ch, ch1, prime)
		ch = ch1
	}
}

func generate(ch chan<- int){
	for i := 2; ; i++{
		fmt.Printf(" i: %d\n", i)
		ch <- i
	}
}

func filter(in <-chan int, out chan<- int, prime int){
	for {
		i := <- in
		fmt.Printf("gid: %d, index: %d, prime: %d  \t",goid.Get(), i, prime)
		if i % prime != 0 {

			out <- i
			fmt.Println()
		}
	}
}
