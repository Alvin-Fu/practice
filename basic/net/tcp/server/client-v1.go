package server

import (
	"bufio"
	"net"
	"practice/basic/net/tcp/client"
	"practice/basic/net/tcp/protocol"
	"strconv"
	"strings"
	"sync"
	"sync/atomic"
	"time"
)

//存储用户临时信息
type ClientV1 struct {
	seq int64

	net.Conn            //tcp连接信息
	ConnTime  time.Time //建立连接时间
	LoginTime time.Time //登录时间
	state     int32     // 0

	writeLock sync.Mutex
	Writer    *bufio.Writer
	Reader    *bufio.Reader

	uid      int64 //用户id，只有完成登录后才不是0
	ip       string
	port     int32
	reqCount uint64
	// 登录信息
	baseInfo atomic.Value

	exitChan chan struct{}
	kickChan chan int32

	heartbeatTime time.Duration
}

//新建一个空白session
func NewClient(c net.Conn) *ClientV1 {
	s := &ClientV1{
		seq:           0,
		Conn:          c,
		ConnTime:      time.Now(),
		Reader:        bufio.NewReaderSize(c, protocol.DefaultReadeBufSize),
		Writer:        bufio.NewWriterSize(c, protocol.DefaultWriteBufSize),
		state:         int32(client.ClientStatusConnect),
		kickChan:      make(chan int32),
		exitChan:      make(chan struct{}),
		heartbeatTime: 20 * time.Second,
	}
	ips := strings.Split(c.RemoteAddr().String(), ":")
	if len(ips) >= 2 {
		s.ip = net.ParseIP(ips[0]).String()
		p, _ := strconv.ParseInt(ips[1], 10, 64)
		s.port = int32(p)
	}
	return s
}
