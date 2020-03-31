package main

import (
	"net"
	"log"
	"time"
	"fmt"
	"encoding/binary"
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
		body := []byte("hello")
		data := make([]byte, 12)
		data[0] = byte('v')
		data[1] = byte('1')
		binary.BigEndian.PutUint32(data[2:6], uint32(12))
		data[6] = byte(1)
		copy(data[protocol.HeadTotalLength:], body)
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