package main

import (
	client2 "practice/basic/net/tcp/client"
)

func main() {
	client := client2.NewClient()
	client.Connect("tcp", "127.0.0.1:8080")
	for i := 0; i < 30; i++ {
		client.Send([]byte("hello"))
		//time.Sleep(1 * time.Second)
	}
	//exit := make(chan struct{})
	//<-exit
}
