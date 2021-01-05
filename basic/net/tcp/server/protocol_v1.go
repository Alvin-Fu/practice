package server

import (
	"fmt"
	"net"
	"practice/basic/net/tcp/protocol"
)

type protocolV1 struct {
	reqPack  *protocol.Packet
	respPack *protocol.Packet
	dec      *protocol.Decoder
	enc      *protocol.Encoder
}

func (p *protocolV1) IOLoop(conn net.Conn) error {
	client := NewClient(conn)
	p.dec = protocol.NewDecoder(client.Reader)
	p.enc = protocol.NewEncoder(client.Writer)
	for {
		err := p.readPacket(client)
		if err != nil {
			fmt.Println(err)
			break
		}
	}
	return nil
}

func (p *protocolV1) readPacket(client *ClientV1) error {
	err := p.dec.Decode(p.reqPack)
	if err != nil {
		return err
	}
	fmt.Println(*p.reqPack)
	return nil
}
