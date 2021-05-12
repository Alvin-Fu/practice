package server

import (
	"fmt"
	"log"
	"net"
	//"practice/basic/net/udp/protocol"
	"practice/basic/net/udp/protocol"
	"practice/lib/util"
)

type ServerV1 struct {
	conn *net.UDPConn
}

func NewServerV1(addr string) *ServerV1 {
	udpAddr, err := net.ResolveUDPAddr("udp", addr)
	if err != nil {
		log.Fatalf("udp resolve err: %v", err)
	}
	conn, err := net.ListenUDP("udp", udpAddr)
	if err != nil {
		log.Fatalf("udp listen err: %v", err)
	}
	return &ServerV1{
		conn: conn,
	}
}

func (s *ServerV1) Start() {
	pack := protocol.NewPacket(protocol.PacketVersionV1)
	data := make([]byte, 100000)
	wg := util.WaitGroupWrapper{}
	for {
		_, _, err := s.conn.ReadFromUDP(data)
		if err != nil {
			fmt.Println(err)
			continue
		}
		err = pack.ParseHead(data[:7])
		if err != nil {
			continue
		}
		wg.Wrap(func() {
			s.Handle(data[7:pack.Length()])
		})
	}
	wg.Wait()
}

func (s *ServerV1) Handle(data []byte) {
	fmt.Printf("data: %s", string(data))
}
