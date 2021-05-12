package protocol

import (
	"net"
	"io"
	"encoding/binary"
	"log"
)

const  (
	VersionLength = 2
	HeadTotalLength = 7
)

type Protocol interface {
	IOLoop(conn net.Conn) error
}

func SendFrameResp(w io.Writer, frameType int32,  resp []byte,)(int, error){
	enBuf := make([]byte, 5)
	binary.BigEndian.PutUint32(enBuf, uint32(len(resp)))
	enBuf[4] = byte(frameType)
	n, err := w.Write(enBuf)
	if err != nil {
		log.Fatalf("write err: %v", err)
		return 0, err
	}
	n, err = w.Write(resp)
	return n + 5, err
}


