package server

import (
	"byrpc/sapi/putil/log"
	"net"
	"practice/lib/util"
	"practice/tool/store-performance/conf"
	"practice/tool/store-performance/server/http"
	"practice/tool/store-performance/service"
	"sync/atomic"
)

type StoreSvr struct {
	opts         atomic.Value
	storeService *service.StoreService
	httpSvr      *http.HTTPServer
	waitGroup    util.WaitGroupWrapper
	exitChan     chan struct{}
}

func NewStoreSvr(opts *conf.Options) *StoreSvr {
	s := &StoreSvr{
		exitChan:     make(chan struct{}),
		storeService: service.NewStoreService(),
	}
	s.swapOpts(opts)
	return s
}

func (s *StoreSvr) getOpts() *conf.Options {
	return s.opts.Load().(*conf.Options)
}

func (s *StoreSvr) swapOpts(opts *conf.Options) {
	s.opts.Store(opts)
}

func (s *StoreSvr) Main() error {
	opts := s.getOpts()
	if opts.HTTPAddr != "" {
		httpListener, err := net.Listen("tcp", s.getOpts().HTTPAddr)
		if err != nil {
			plog.Fatalf("listen (%s) failed - %s", s.getOpts().HTTPAddr, err)
			return err
		}
		httpServer := http.NewHTTPSvr(httpListener, s.storeService)
		if httpServer != nil {
			s.waitGroup.Wrap(func() {
				httpServer.Start(s.exitChan)
			})
		}
		s.httpSvr = httpServer
	}
	return nil
}

func (s *StoreSvr) Exit() error {
	close(s.exitChan)
	return nil
}
