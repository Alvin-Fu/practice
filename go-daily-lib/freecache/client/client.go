package main

import (
	"net"
	"log"
	"time"
	"fmt"
		"practice/go-daily-lib/freecache/protocol"
)

func init(){
	log.SetFlags( log.Lshortfile |log.LstdFlags| log.Llongfile)
}
var exitChan = make(chan struct{})
func main(){
	conn, err := net.Dial("tcp", "127.0.0.1:8080")
	if err != nil {
		log.Fatalf("dial err: %v", err)
		return
	}
	go onMessage(conn)
	<- exitChan
	conn.Close()
}

func onMessage(conn net.Conn) {
	for {
		packet := new(protocol.Packet)
		packet.Fill([2]byte{'v', '1'}, '1', []byte("hello"))
		data, err := packet.Marshal()
		if err != nil {
			log.Fatalf("packet marshal err: %v", err)
			exitChan <- struct{}{}
			break
		}
		n, err := conn.Write(data)
		if err != nil {
			log.Fatalf("write err: %v", err)
			exitChan <- struct{}{}
			break
		}
		fmt.Println(n)
		time.Sleep(2 * time.Second)
	}
}

