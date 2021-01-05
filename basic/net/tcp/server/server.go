package server

import (
	"net"
	"practice/basic/net/tcp/protocol"
	"practice/lib/util"
	"runtime"
	"strings"
)

type ServerV1 struct {
	listener net.Listener
}

func NewServer(network, addr string) *ServerV1 {
	listener, err := net.Listen(network, addr)
	if err != nil {
		return nil
	}
	return &ServerV1{
		listener: listener,
	}
}

func (s *ServerV1) Start() {
	waitGroup := util.WaitGroupWrapper{}
	for {
		conn, err := s.listener.Accept()
		if err != nil {
			if nerr, ok := err.(net.Error); ok && nerr.Temporary() {
				runtime.Gosched()
				continue
			}
			// theres no direct way to detect this error because it is not exposed
			if !strings.Contains(err.Error(), "use of closed network connection") {

			}
			break
		}
		waitGroup.Wrap(func() {
			s.readLoop(conn)
		})
	}
	waitGroup.Wait()
}

func (s *ServerV1) readLoop(conn net.Conn) {
	prot := &protocolV1{
		reqPack: protocol.NewPacket(protocol.PacketVersionV1),
	}
	prot.IOLoop(conn)
}
