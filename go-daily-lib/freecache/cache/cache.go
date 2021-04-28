package cache

import (
	"net"
	"practice/go-daily-lib/freecache/uitl"
	"sync"
	"sync/atomic"
	"time"
		"os"
	"practice/go-daily-lib/freecache/protocol"
	"log"
	"github.com/coocood/freecache"
)

const (
	SvrStatusInvalid int32 = -1
	SvrStatusInit    int32 = 0
	SvrStatusRunning int32 = 1
	SvrStatusExiting int32 = 2
)

var (
	CacheSize = 256 * 1024 * 1024
)

var Cache *CacheSvr

type CacheSvr struct {
	clientIdReq int64

	caches *freecache.Cache


	sync.Mutex
	opts      atomic.Value
	startTime time.Time
	status    int32
	exitChan  chan struct{}
	waitGroup uitl.WaitGroupWrapper

	tcpListener net.Listener

	clientMtx sync.RWMutex
	clients map[int64]*clientV1
}

func NewCacheSvr(opts *Options) *CacheSvr {
	cacheSvr := &CacheSvr{
		startTime: time.Now(),
		exitChan:  make(chan struct{}),
		status: SvrStatusInit,
		clients: make(map[int64]*clientV1),
		caches: freecache.NewCache(CacheSize),
	}
	cacheSvr.swapOpts(opts)

	return cacheSvr
}

func (cs *CacheSvr) Main() {
	var err error
	cs.tcpListener, err = net.Listen("tcp", cs.getOpts().TCPAddr)
	if err != nil {
		log.Printf("net listen err: %v", err)
		os.Exit(1)
	}
	ctx := new(context)
	ctx.cs = cs
	tcpSever := &tcpSever{ctx}
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

func (cs *CacheSvr) AddClient(clientId int64, client *clientV1){
	cs.clientMtx.Lock()
	defer cs.clientMtx.Unlock()
	cs.clients[clientId] = client
}
func (ca *CacheSvr) RemoveClient(clientId int64){
	ca.clientMtx.Lock()
	defer ca.clientMtx.Unlock()
	delete(ca.clients, clientId)
}

func (cs *CacheSvr) swapOpts(opts *Options){
	cs.opts.Store(opts)
}

func (cs *CacheSvr) getOpts()*Options{
	return cs.opts.Load().(*Options)
}