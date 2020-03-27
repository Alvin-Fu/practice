package cache

import (
	"net"
	"practice/go-daily-lib/freecache/uitl"
	"sync"
	"sync/atomic"
	"time"
		"os"
	"practice/go-daily-lib/freecache/protocol"
	"practice/go-daily-lib/freecache/log"
)

const (
	SvrStatusInvalid int32 = -1
	SvrStatusInit    int32 = 0
	SvrStatusRunning int32 = 1
	SvrStatusExiting int32 = 2
)

type CacheSvr struct {
	sync.Mutex
	opts      atomic.Value
	startTime time.Time
	status    int32
	exitChan  chan struct{}
	waitGroup uitl.WaitGroupWrapper

	tcpListener net.Listener
}

func NewCacheSvr(opts *Options) *CacheSvr {
	cacheSvr := &CacheSvr{
		startTime: time.Now(),
		exitChan:  make(chan struct{}),
		status: SvrStatusInit,
	}
	cacheSvr.swapOpts(opts)

	return cacheSvr
}

func (cs *CacheSvr) Main() {
	var err error
	cs.tcpListener, err = net.Listen("tcp", cs.getOpts().TCPAddr)
	if err != nil {
		log.Debugf("net listen err: %v", err)
		os.Exit(1)
	}
	tcpSever := &tcpSever{}
	cs.waitGroup.Wrap(func() {
		protocol.TCPSever(cs.tcpListener, tcpSever)
	})
	cs.SetSvrStatus(SvrStatusRunning)
}

func (cs *CacheSvr) Exit() {}

func (cs *CacheSvr) Status()int32{
	return cs.status
}

func (cs *CacheSvr) SetSvrStatus(status int32){
	cs.status = status
}

func (cs *CacheSvr) swapOpts(opts *Options){
	cs.opts.Store(opts)
}

func (cs *CacheSvr) getOpts()*Options{
	return cs.opts.Load().(*Options)
}