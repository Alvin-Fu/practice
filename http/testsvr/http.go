package testsvr

import (
	"context"
	"fmt"
	"github.com/gorilla/mux"
	"github.com/pkg/errors"
	plog "log"
	"net"
	"net/http"
	"net/http/pprof"
	"practice/http/util"
)

type HTTPSever struct {
	listener net.Listener
	router   *mux.Router
	//ctx       *context.Context
	server    *http.Server
	handler   interface{}
	waitGroup util.WaitGroupWrapper
}

func (s *HTTPSever) Start(exitChan <-chan struct{}) {

	s.server = &http.Server{Handler: removeTrailingSlash(s.router)}
	s.waitGroup.Wrap(func() {
		s.server.Serve(s.listener)
	})
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
		plog.Fatalf("Http sever shut down fail err: %v", err)
	}
	//s.waitGroup.Wait()
	return
}
func (s *HTTPSever) index(w http.ResponseWriter, r *http.Request) {
	fmt.Println("hello")
	fmt.Fprintf(w, "%s\n", "hello")
}

func (s *HTTPSever) RegRouter() error {
	if s.router == nil {
		plog.Fatalf("Invalid router!")
		return errors.Errorf("Invalid router!")
	}
	s.router.HandleFunc("/", s.index)
	s.router.HandleFunc("/debug/pprof/", pprof.Index)
	s.router.HandleFunc("/debug/pprof/cmdline", pprof.Cmdline)
	s.router.HandleFunc("/debug/pprof/profile", pprof.Profile)
	s.router.HandleFunc("/debug/pprof/symbol", pprof.Symbol)
	s.router.HandleFunc("/debug/pprof/trace", pprof.Trace)
	s.router.HandleFunc("/debug/pprof/{name:[a-z]+}", pprofHelper)
	sr := s.router.PathPrefix("/v1/").Subrouter()
	handler := &HTTPHandlerV1{}
	s.handler = handler
	sr.HandleFunc("/matchs", handler.GetMatchInfo).Methods("GET")
	sr.HandleFunc("/matchs/{mid:[1-9][0-9]*}", handler.GetOneMatchInfo).Methods("GET")
	return nil
}

func NewHTTPSever(listener net.Listener) *HTTPSever {
	svr := &HTTPSever{listener: listener, router: mux.NewRouter()}
	if err := svr.RegRouter(); err != nil {
		fmt.Printf("Router register fail err: %v\n", err)
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
		//fmt.Println(r.URL.Path)
		////r.URL.Path = strings.TrimPrefix(r.URL.Path, "/")
		//fmt.Println(r.URL.Path)
	handle:
		next.ServeHTTP(w, r)
	})
}
func pprofHelper(w http.ResponseWriter, r *http.Request) {
	name := mux.Vars(r)["name"]
	pprof.Handler(name).ServeHTTP(w, r)
}
