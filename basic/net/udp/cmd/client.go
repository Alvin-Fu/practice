package main

import (
	"fmt"
	"practice/basic/net/udp/client"
)

func main() {
	cl := client.NewClientV1("120.79.93.211:8080")
	data := make([]byte, 0)
	for i := 0; i < 65500; i++ {
		data = append(data, 'a')
	}
	fmt.Println(len(data))
	fmt.Println(cl.Send(data))
}
