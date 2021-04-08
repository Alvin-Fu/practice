package http

import (
	"byrpc/sapi/putil/log"
	goctx "context"
	"domino/lib/util"
	"encoding/json"
	"fmt"
	"github.com/gorilla/mux"
	"net"
	"net/http"
	"net/http/pprof"
	"strings"
)

// HTTPServer is a instance of admin service.
// http://api.boyaa.robot:8080
type HTTPServer struct {
	listener  net.Listener
	router    *mux.Router
	server    *http.Server
	waitGroup util.WaitGroupWrapper
}

// Start use to start http server.
func (s *HTTPServer) Start(exitChan <-chan struct{}) error {
	plog.Debugf("http server start")
	s.server = &http.Server{Handler: removeTrailingSlash("/debug/pprof/", s.router)}
	s.server.Serve(s.listener)

	// wait exit signal
	for {
		select {
		case <-exitChan:
			goto exit
		}
	}
exit:
	plog.Warnf("http svr stop")
	s.Stop()
	return nil
}

// Stop use to stop http server.
func (s *HTTPServer) Stop() error {
	err := s.server.Shutdown(goctx.TODO())
	if err != nil {
		plog.Fatalf("http server shutdown fail: %s", err)
	}
	s.waitGroup.Wait()
	plog.Debugf("http server exit")
	return err
}

func (s *HTTPServer) index(w http.ResponseWriter, r *http.Request) {
	//resp := s.svr.State()
	//resp = append(resp, '\n')
	//w.Write(resp)
}

// regRouter register handler for "api.boyaa.robot"
func (s *HTTPServer) regRouter() error {
	if s.router == nil {
		return fmt.Errorf("invalid nil router")
	}
	s.router.HandleFunc("/", s.index)

	s.router.HandleFunc("/debug/pprof/", pprof.Index)
	s.router.HandleFunc("/debug/pprof/cmdline", pprof.Cmdline)
	s.router.HandleFunc("/debug/pprof/profile", pprof.Profile)
	s.router.HandleFunc("/debug/pprof/symbol", pprof.Symbol)
	s.router.HandleFunc("/debug/pprof/trace", pprof.Trace)
	s.router.HandleFunc("/debug/pprof/{name:[a-z]+}", pprofHelper)

	// router for v1
	return nil
}

// 返回一个路由，用于注册业务自己的路由规则
func (s *HTTPServer) Router() *mux.Router {
	return s.router
}

// response make response in v1 and send to client.
func Response(w http.ResponseWriter, statusCode int, errmsg string, respJSON []byte) error {
	// use statusCode as error code
	w.Header().Add("Content-Type", "application/json; charset=utf-8")
	w.WriteHeader(statusCode)

	resp, _ := json.Marshal(struct {
		Errmsg string
		Result string
	}{errmsg, string(respJSON)})

	resp = append(resp, '\n')
	_, err := w.Write(resp)
	return err
}

// NewHTTPSvr create a http server instance.
func NewHTTPSvr(listener net.Listener) *HTTPServer {
	svr := &HTTPServer{listener: listener, router: mux.NewRouter()}
	if err := svr.regRouter(); err != nil {
		plog.Fatalf("http server register router fail: %s", err)
		return nil
	}
	return svr
}

// removeTrailingSlash is helper function to remove trailing '/',
// note we don't use mux.NewRouter().StrictSlash(true),
// since this will not only redirect '/foo/' to '/foo',
// but may also redirect POST to GET.
// Path after replace will like: '^/.*[^/]$'
func removeTrailingSlash(excludePrefix string, next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if len(r.URL.Path) == 0 || r.URL.Path == "/" {
			r.URL.Path = "/"
			goto handle
		} else if strings.HasPrefix(r.URL.Path, excludePrefix) {
			goto handle
		}

		// only remove last trailing slash
		r.URL.Path = strings.TrimSuffix(r.URL.Path, "/")
	handle:
		next.ServeHTTP(w, r)
	})
}

func pprofHelper(w http.ResponseWriter, r *http.Request) {
	name := mux.Vars(r)["name"]
	pprof.Handler(name).ServeHTTP(w, r)
}
