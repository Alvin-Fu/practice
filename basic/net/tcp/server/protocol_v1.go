package server

import (
	"fmt"
	"net"
	"practice/basic/net/tcp/protocol"
	"time"
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
	go func() {
		for i := 0; i < 1000; i++ {
			err := p.writePacket(client)
			if err != nil {
				fmt.Println("write packet err", err)
				break
			}
			time.Sleep(time.Second * 20)
		}
		time.Sleep(time.Minute * 2)
		err := p.writePacket(client)
		if err != nil {
			fmt.Println("write packet err", err)
		}
		return
	}()
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

func (p *protocolV1) writePacket(client *ClientV1) error {
	var packet = protocol.NewPacket(protocol.PacketVersionV1)
	packet.Fill(protocol.PacketVersionV1, 1, []byte("1"))
	err := p.enc.Encode(packet)
	if err != nil {
		fmt.Println(err)
		return err
	}
	return client.Writer.Flush()
	return nil
}
