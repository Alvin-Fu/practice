package main

import "practice/basic/net/udp/client"

func main() {
	cl := client.NewClientV1("127.0.0.1:8080")
	data := make([]byte, 2000)
	data[1] = 'a'
	cl.Send(data)
}
