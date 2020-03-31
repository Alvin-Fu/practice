package cache

import (
	"net"
	"sync"
		"bufio"
	"log"
	"time"
)
const defaultBufferSize =8 * 1024

const (
	statusInit = iota
	statusDisconnected
	statusConnected
	statusLogined
	statusIdentified
	statusClosing
)

type clientV1 struct {
	ID int64
	ctx *context
	net.Conn

	writeLock sync.Mutex

	status int32

	Reader *bufio.Reader
	Writer *bufio.Writer

	connectTime time.Time

	HeartbeatInterval time.Duration
}

func NewClient(id int64, conn net.Conn, ctx *context)*clientV1 {
	return &clientV1{
		ID: id,
		ctx: ctx,
		Conn: conn,
		status: statusInit,

		Reader: bufio.NewReaderSize(conn, defaultBufferSize),
		Writer: bufio.NewWriterSize(conn, 0),
		connectTime: time.Now(),

		HeartbeatInterval: 2,
	}
}

func (c *clientV1)Flush()error{
	err := c.Writer.Flush()
	if err != nil {
		log.Printf("writer flush err: %v", err)
		return err
	}
	return nil
}

