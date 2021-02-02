package client

import (
	"fmt"
	"log"
	"net"
	"practice/basic/net/udp/protocol"
)

type ClientV1 struct {
	conn *net.UDPConn
	pack *protocol.Packet
}

func NewClientV1(addr string) *ClientV1 {
	udpAddr, err := net.ResolveUDPAddr("udp", addr)
	if err != nil {
		log.Fatalf("udp resolve err: %v", err)
	}
	conn, err := net.DialUDP("udp", nil, udpAddr)
	if err != nil {
		log.Fatalf("udp listen err: %v", err)
	}
	return &ClientV1{
		conn: conn,
		pack: protocol.NewPacket(protocol.PacketVersionV1),
	}
}

func (c *ClientV1) Send(data []byte) error {
	c.pack.Fill(protocol.PacketVersionV1, 1, data)
	payload, _ := c.pack.Marshal()
	n, err := c.conn.Write(payload)
	if err != nil {
		return err
	}
	fmt.Println(n)
	return nil
}
