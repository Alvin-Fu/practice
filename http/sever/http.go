package sever

import (
	"context"
	"fmt"
	"github.com/gorilla/mux"
	"github.com/pkg/errors"
	"net"
	"net/http"
	_ "net/http/pprof"
	"strings"
)

type HTTPSever struct {
	listener net.Listener
	router   *mux.Router
	ctx      *context.Context
	server   *http.Server
	handler  interface{}
}

func (s *HTTPSever) Start(exitChan <-chan struct{}) {

	s.server = &http.Server{Handler: removeTrailingSlash(s.router)}
	//s.waitGroup.Wrap(func() {
	s.server.Serve(s.listener)
	//})
	fmt.Println("Http sever start success!")
	// Todo: 监听信号，到收到时停止web服务
	for {
		select {
		case <-exitChan:
			goto exit
		}
	}
exit:
	s.Stop()
	return
}

func (s *HTTPSever) Stop() {
	// 优雅的关闭连接
	err := s.server.Shutdown(context.TODO())
	if err != nil {
		fmt.Println("Http sever shut down fail err: %v", err)
	}
	//s.waitGroup.Wait()
	fmt.Println("Http sever shut down success!")
	return
}
func (s *HTTPSever) index(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "%s\n", "hello")
}

func (s *HTTPSever) RegRouter() error {
	if s.router == nil {
		fmt.Println("Invalid router!")
		return errors.Errorf("Invalid router!")
	}
	s.router.HandleFunc("/", s.index)

	handler := &HTTPHandlerV1{ctx: s.ctx}
	//s.handler = handler

	sr := s.router.PathPrefix("/v/").Subrouter()
	sr.HandleFunc("/match", handler.GetMatchInfo).Methods("GET")
	sr.HandleFunc("/match/{mid:[1-9][0-9]*}", handler.GetOneMatchInfo).Methods("GET")
	return nil
}

func NewHTTPSever(listener net.Listener) *HTTPSever {
	svr := &HTTPSever{listener: listener, router: mux.NewRouter()}
	if err := svr.RegRouter(); err != nil {
		fmt.Println("Router register fail err: %v", err)
		return nil
	}
	return svr
}

func removeTrailingSlash(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if len(r.URL.Path) == 0 || r.URL.Path == "/" {
			r.URL.Path = "/"
			goto handle
		}
		r.URL.Path = strings.TrimPrefix(r.URL.Path, "/")
	handle:
		next.ServeHTTP(w, r)
	})
}
