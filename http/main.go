package main

import (
	"fmt"
	"net"
	"practice/httptice/http/sever"
	"sync"
)

func main() {
	var ch = make(chan struct{})
	tcpListener, err := net.Listen("tcp", "127.0.0.1:8088")
	if err != nil {
		fmt.Println("Tcp connect fail err: %v", err)
		return
	}
	httpSever := sever.NewHTTPSever(tcpListener)
	wg := sync.WaitGroup{}
	wg.Add(1)
	go func() {
		httpSever.Start(ch)
		wg.Done()
	}()
	wg.Wait()

}
