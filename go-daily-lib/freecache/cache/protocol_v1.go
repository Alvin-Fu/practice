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
	"practice/go-daily-lib/freecache/pb"
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
		req, err := p.readRequest(client, 0)
		if err != nil {
			if err == io.EOF{
				err = nil
			} else {
				err = fmt.Errorf("failed to read command - %v", err)
			}
			ioErr = err
			log.Printf("read request err: %v", err)
			break
		}
		if  req == nil {
			// 心跳包
			continue
		}
		// 根据包内容进行操作
		log.Printf("req: %s", req.String())
		go p.Exec(client, req, p.execCallBack)
	}
	if ioErr != nil {
		log.Fatalf("ioerr: %v", ioErr)
		return ioErr
	}
	log.Println(ioErr)
	return nil
}


func (p *protocolV1) Exec(client *clientV1, req *PbCache.CacheReq, cb execCallBack){
	var err error
	var resp []byte
	switch req.GetOptionType() {
	case PbCache.OptionType_Option_Type_Get:
		// 查询
		value, err := Cache.caches.Get([]byte(req.GetKey()))
		if err != nil {
			log.Fatalf("caches get err: %v", err)
			resp, _ = p.MakeCacheResponse(err.Error(), protocol.CacheSvrErrCode, "")
		} else{
			resp, _ = p.MakeCacheResponse("", protocol.CacheSvrOK, string(value))
		}

	case PbCache.OptionType_Option_Type_Set:
		// 设置
		err = Cache.caches.Set([]byte(req.GetKey()), []byte(req.GetValue()), int(req.GetExpire()))
		if err != nil {
			log.Fatalf("caches get err: %v", err)
			resp, _ = p.MakeCacheResponse(err.Error(), protocol.CacheSvrErrCode, "")
		} else{
			resp, _ = p.MakeCacheResponse("", protocol.CacheSvrOK, "")
		}
	case PbCache.OptionType_Option_Type_Del:
		// 删除
		if Cache.caches.Del([]byte(req.GetKey())){
			resp, _ = p.MakeCacheResponse("", protocol.CacheSvrOK, "ok")
		}
	default:


	}
	cb(client, req, resp, err)
}

func (p *protocolV1) Send(client *clientV1, data []byte) error{
	client.writeLock.Lock()
	defer client.writeLock.Unlock()
	n, err := p.SendResponse(client.Writer, 0, data)
	if err != nil {
		log.Fatalf("send response err: %v", err)
		return err
	}
	log.Printf("send count: %d", n)
	client.Flush()
	return nil
}


func (p *protocolV1) SendResponse(w io.Writer, frameType int32, payload []byte)(int, error){
	return protocol.SendFrameResp(w, frameType, payload)
}

func (p *protocolV1) MakeCacheResponse(errMsg string, errCode int32, rue string)([]byte, error){
	resp := new(PbCache.CacheResp)
	resp.ErrMsg = &errMsg
	resp.ErrCode = &errCode
	resp.Rue = &rue
	return resp.Marshal()
}


func (p *protocolV1) execCallBack(client *clientV1, req interface{}, resp []byte, err error){
	if err != nil {
		return
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

func (p *protocolV1) readRequest(client *clientV1, skit int)(*PbCache.CacheReq, error){
	headBuf := make([]byte, protocol.HeadTotalLength)
	_, err := io.ReadFull(client.Reader, headBuf)
	if err != nil {
		//log.Fatalf("read full err: %v", err)
		return nil, err
	}
	// 使用pb或者json
	bodyLen := int32(binary.BigEndian.Uint32(headBuf[2:6])) - int32(protocol.HeadTotalLength)
	body := make([]byte, bodyLen)
	_, err = io.ReadFull(client.Reader, body)
	if err != nil {
		//log.Fatalf("read full err: %v", err)
		return nil, err
	}
	req := new(PbCache.CacheReq)
	err = req.Unmarshal(body)
	if err != nil {
		log.Fatalf("unmarshal err: %v", err)
		return nil, err
	}
	if req.GetKey() == ""{
		log.Fatalf("nill req")
		return nil, fmt.Errorf("nil req")
	}
	return req, nil
}

func (p *protocolV1) processRequest(client *clientV1){
	//Todo: 对请求进行处理

}

func (p *protocolV1) kick(){}

func (p *protocolV1) logout(client *clientV1){}
