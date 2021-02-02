package main

import (
	"fmt"
	"practice/basic/net/udp/server"
)

func main() {
	svr := server.NewServerV1("127.0.0.1:8080")
	fmt.Println("hello")
	svr.Start()
}
