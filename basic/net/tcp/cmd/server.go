package main

import "practice/basic/net/tcp/server"

func main() {
	svr := server.NewServer("tcp", "0.0.0.0:8080")
	svr.Start()
}
