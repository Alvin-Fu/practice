package main

import (
	"net"
		"time"
)

func main(){
	for i := 0; i < 100; i ++{
		conn, err := net.Dial("tcp", "127.0.0.1:5000")
		if err != nil {
			panic(err)
		}
		conn.Write([]byte("ok"))
		time.Sleep(100 * time.Millisecond)
	}

}
