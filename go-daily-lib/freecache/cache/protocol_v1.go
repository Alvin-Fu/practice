package cache

import (
	"net"
	"io"
	"practice/go-daily-lib/freecache/log"
)

type protocolV1 struct {

}

func (p *protocolV1) IOLoop(conn net.Conn)error{
	client := NewClient()
	for {
		//
		p.readRequest(client, 0)
	}
	return nil
}

func (p *protocolV1) readRequest(client *clientV1, skit int)([]byte, error){
	body := make([]byte, 0)
	_, err := io.ReadFull(client.Reader, body)
	if err != nil {
		log.Fatalf("read full err: %v", err)
		return nil, err
	}
	// 使用pb或者json
	return body, nil
}

func (p *protocolV1) processRequest(client *clientV1){
	//Todo: 对请求进行处理

}

func (p *protocolV1) Send(client *clientV1, data []byte) {}
