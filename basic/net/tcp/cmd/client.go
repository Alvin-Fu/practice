package main

import (
	"fmt"
	client2 "practice/basic/net/tcp/client"
	"time"
)

func main() {
	client := client2.NewClient()
	err := client.Connect("tcp", "192.168.28.176:8080")
	if err != nil {
		fmt.Println(err)
		return
	}
	for i := 0; i < 30; i++ {
		client.Send([]byte("hello"))
		time.Sleep(1 * time.Second)
	}
	//exit := make(chan struct{})
	//<-exit
}
