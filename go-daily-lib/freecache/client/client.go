package main

import (
	"net"
	"log"
		"fmt"
		"practice/go-daily-lib/freecache/protocol"
	"practice/go-daily-lib/freecache/pb"
	"github.com/gogo/protobuf/proto"
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
		req := new(PbCache.CacheReq)
		req.OptionType = PbCache.OptionType_Option_Type_Set.Enum()
		req.Key = proto.String("hello")
		req.Value = proto.String("world")
		d, _ := req.Marshal()
		packet := new(protocol.Packet)
		packet.Fill([2]byte{'v', '1'}, '1', d)
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
		exitChan <- struct{}{}
		break
	}
}

