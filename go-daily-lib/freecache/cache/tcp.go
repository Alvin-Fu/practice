package cache

import (
	"net"
			"practice/go-daily-lib/freecache/protocol"
	"log"
	"time"
	"math/rand"
)

type tcpSever struct {
	ctx *context
}

func (s *tcpSever) Handle( conn net.Conn){
	// 设置读超时
	t := time.Duration(rand.Intn(3)+1) * s.ctx.cs.getOpts().FirstReadTimeOut
	log.Printf("read timeout: %d", t)
	conn.SetReadDeadline(time.Now().Add(t * time.Second))
	var prot protocol.Protocol
	prot = &protocolV1{s.ctx}
	err := prot.IOLoop(conn)
	if err != nil {
		//log.Fatalf("io loop err: %v", err)
		return
	}

}
