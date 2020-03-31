package cache

import (
	"net"
	"io"
	"log"
	"sync/atomic"
	"fmt"
	"practice/go-daily-lib/freecache/protocol"
	"encoding/binary"
	"time"
)

type execCallBack func(client *clientV1, req interface{}, resp []byte, err error)

type protocolV1 struct {
	ctx *context
}

func (p *protocolV1) IOLoop(conn net.Conn)error{
	var ioErr error
	clientId := atomic.AddInt64(&p.ctx.cs.clientIdReq, 1)
	client := NewClient(clientId, conn, p.ctx)
	p.ctx.cs.AddClient(clientId, client)
	for {
		//
		if client.HeartbeatInterval > 0 {
			client.SetReadDeadline(time.Now().Add(client.HeartbeatInterval * 2 * time.Second))
		} else {

		}
		data, err := p.readRequest(client, 0)
		if err != nil {
			if err == io.EOF{
				err = nil
			} else {
				err = fmt.Errorf("")
			}
			ioErr = err
			break
		}
		if len(data) == 0 {
			// 心跳包
			continue
		}
		// 根据包内容进行操作
		log.Printf("data: %s", string(data))
		p.Exec(client, data, p.execCallBack)
	}
	return ioErr
}


func (p *protocolV1) Exec(client *clientV1, data []byte, cb execCallBack){

}

func (p *protocolV1) Send(client *clientV1, data []byte) error{
	client.writeLock.Lock()
	defer client.writeLock.Unlock()
	_, err := p.SendResponse(client.Writer, 0, data)
	if err != nil {
		log.Fatalf("send response err: %v", err)
		return err
	}
	client.Flush()
	return nil
}

func (p *protocolV1) SendResponse(w io.Writer, frameType int32, payload []byte)(int, error){
	return protocol.SendFrameResp(w, frameType, payload)
}

func (p *protocolV1) execCallBack(client *clientV1, req interface{}, resp []byte, err error){
	if err != nil {

	}
	if resp != nil {
		sErr := p.Send(client, resp)
		if sErr != nil {
			log.Printf("send err: %v", err)
			goto kickClient
		}
	}
	// 断开连接
	kickClient:
		p.kick()
}

func (p *protocolV1) readRequest(client *clientV1, skit int)([]byte, error){
	headBuf := make([]byte, protocol.HeadTotalLength)
	n, err := io.ReadFull(client.Reader, headBuf)
	if err != nil {
		log.Fatalf("read full err: %v, n: %d", err, n)
		return nil, err
	}
	// 使用pb或者json
	bodyLen := int32(binary.BigEndian.Uint32(headBuf[2:6])) - int32(protocol.HeadTotalLength)
	body := make([]byte, bodyLen)
	_, err = io.ReadFull(client.Reader, body)
	if err != nil {
		log.Fatalf("read full err: %v, n: %d", err, n)
		return nil, err
	}
	return body, nil
}

func (p *protocolV1) processRequest(client *clientV1){
	//Todo: 对请求进行处理

}

func (p *protocolV1) kick(){}
