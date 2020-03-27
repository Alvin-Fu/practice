package cache

import (
	"net"
	"sync"
		"bufio"
	"practice/go-daily-lib/freecache/log"
)

type clientV1 struct {
	ID int64
	net.Conn

	writeLock sync.Mutex

	status int32

	Reader *bufio.Reader
	Writer *bufio.Writer
}

func NewClient()*clientV1 {
	return &clientV1{}
}

func (c *clientV1)Flush()error{
	err := c.Writer.Flush()
	if err != nil {
		log.Debugf("writer flush err: %v", err)
		return err
	}
	return nil
}

