package client

import (
	"bufio"
	"fmt"
	"net"
	"practice/basic/net/tcp/protocol"
	"sync"
)

type ClientStatus int32

const (
	ClientStatusInit       ClientStatus = 0
	ClientStatusConnect    ClientStatus = 1
	ClientStatusLogin      ClientStatus = 2
	ClientStatusLogout     ClientStatus = 3
	ClientStatusDisconnect ClientStatus = 4
)

func (s ClientStatus) String() string {
	switch s {
	case ClientStatusInit:
		return "init"
	case ClientStatusConnect:
		return "connect"
	case ClientStatusLogin:
		return "login"
	case ClientStatusLogout:
		return "logout"
	case ClientStatusDisconnect:
		return "disconnect"
	default:
		return "init"
	}
}

type ClientV1 struct {
	mtx       sync.RWMutex
	conn      net.Conn
	status    ClientStatus
	reader    *bufio.Reader
	writer    *bufio.Writer
	writeChan chan []byte

	decoder *protocol.Decoder
	encoder *protocol.Encoder

	exitChan chan struct{}
}

func NewClient() *ClientV1 {
	return &ClientV1{
		exitChan: make(chan struct{}),
	}
}

func (c *ClientV1) Status() ClientStatus {
	return c.status
}

func (c *ClientV1) SetStatus(status ClientStatus) {
	c.status = status
}

func (c *ClientV1) Connect(network, addr string) error {
	conn, err := net.Dial(network, addr)
	if err != nil {
		return err
	}
	c.conn = conn
	c.reader = bufio.NewReaderSize(conn, protocol.DefaultReadeBufSize)
	c.writer = bufio.NewWriterSize(conn, protocol.DefaultWriteBufSize)
	c.decoder = protocol.NewDecoder(c.reader)
	c.encoder = protocol.NewEncoder(c.writer)
	c.SetStatus(ClientStatusConnect)
	go c.readLoop()
	c.writeChan = make(chan []byte, 16)
	go c.writeLoop()

	return nil
}

func (c *ClientV1) Disconnect() {
	c.mtx.Lock()
	defer c.mtx.Unlock()
	if c.status != ClientStatusConnect {
		return
	}
	c.SetStatus(ClientStatusDisconnect)
	if c.conn != nil {
		c.conn.Close()
	}
	count := 0
	for {
		select {
		case <-c.writeChan:
			count++
		default:
			goto exit
		}
	}
exit:
	close(c.exitChan)
}

func (c *ClientV1) readLoop() {
	pack := protocol.NewPacket(protocol.PacketVersionV1)
	for {
		err := c.decoder.Decode(pack)
		if err != nil {
			goto exit
		}
		fmt.Println("pack: ", pack)
		select {
		case <-c.exitChan:
			goto exit
		default:
		}
	}
exit:
	c.Disconnect()
}

func (c *ClientV1) Send(data []byte) {
	c.writeChan <- data
}

func (c *ClientV1) writeLoop() {
	pack := protocol.NewPacket(protocol.PacketVersionV1)
	for {
		select {
		case data := <-c.writeChan:
			pack.Fill(protocol.PacketVersionV1, 1, data)
			err := c.encoder.Encode(pack)
			if err != nil {
				goto exit
			}
			//if len(c.writeChan) == 0 {
			err = c.writer.Flush()
			if err != nil {
				goto exit
			}
			//}
		case <-c.exitChan:
			goto exit
		}
	}
exit:
	fmt.Println("exit")
	c.Disconnect()
}
