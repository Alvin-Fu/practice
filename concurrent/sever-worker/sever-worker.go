package main

import (
	"net"
	"time"
	"fmt"
)

func main(){
	l, err := net.Listen("tcp", "127.0.0.1:5000")
	if err != nil {
		panic(err)
	}
	ch := make(chan string)
	go sever(l, ch)
	go pool(ch, 4)
	time.Sleep(100 * time.Second)
}

func sever(l net.Listener, ch chan string){
	for  {
		c, err := l.Accept()
		if err != nil {
			continue
		}
		go handler(c, ch)
	}
}
func pool(ch chan string, n int){
	wch := make(chan int)
	res := make(chan int)
	for i := 0; i < n; i++{
		go logger(wch, res)
	}
	go parse(res)
	for {
		addr := <- ch
		l := len(addr)
		wch <- l
	}
}
func handler(c net.Conn, ch chan string){
	addr := c.RemoteAddr().String()
	ch <- addr
	time.Sleep(100 * time.Millisecond)
	c.Write([]byte("ok"))
	c.Close()
}
func logger(wch chan int, results chan int){
	for  {
		data := <- wch
		data ++
		results <- data
	}
}
func parse(results chan int){
	for   {
		fmt.Println(<- results)
	}
}
