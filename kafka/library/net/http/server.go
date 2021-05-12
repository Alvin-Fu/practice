package http

import (
	goctx "context"
	"log"
	"net"
	"net/http"
	"practice/kafka/producer/option"
	"practice/kafka/util"
)

type Server struct {
	opt       *option.Options
	listener  net.Listener
	waitGroup util.WaitGroupWrapper
	server    *http.Server
}

func NewServer()*Server{}

func (s *Server) Start(exitChan <-chan struct{}) error {


	for {
		select {
		case <-exitChan:
			goto exit
		}
	}
exit:
	s.Stop()
	return nil
}

func (s *Server) Stop() error {
	err := s.server.Shutdown(goctx.TODO())
	if err != nil {
		log.Fatal("http server shutdown fail: %v", err)
	}
	s.waitGroup.Wait()
	return nil
}
