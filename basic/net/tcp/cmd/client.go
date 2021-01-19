package main

import (
	"fmt"
	client2 "practice/basic/net/tcp/client"
	"time"
)

func main() {
	client := client2.NewClient()
	err := client.Connect("tcp", "192.168.81.15:8080")
	if err != nil {
		fmt.Println(err)
		return
	}
	data := make([]byte, 0)
	for i := 0; i < 10; i++ {
		data = append(data, '1')
	}
	for i := 0; i < 100; i++ {
		client.Send(data)
		time.Sleep(1 * time.Second)
	}
	client.Disconnect()
	exit := make(chan struct{})
	<-exit
}
