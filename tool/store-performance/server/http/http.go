package http

import (
	"net"
	"practice/lib/httpsvr"
	"practice/tool/store-performance/service"
)

type HTTPServer struct {
	svr     *http.HTTPServer
	handler interface{}
}

// NewHTTPSvr create a http server instance.
func NewHTTPSvr(listener net.Listener, storeService *service.StoreService) *HTTPServer {
	hsvr := http.NewHTTPSvr(listener)
	if hsvr == nil {
		return nil
	}
	handler := &HTTPHandlerV1{storeService: storeService}
	svr := &HTTPServer{svr: hsvr, handler: handler}
	if err := svr.regRouter(); err != nil {
		return nil
	}
	return svr
}

func (h *HTTPServer) Stop() error {
	return h.svr.Stop()
}

func (h *HTTPServer) Start(exitChan chan struct{}) error {
	return h.svr.Start(exitChan)
}

func (h *HTTPServer) regRouter() error {
	router := h.svr.Router()
	sr := router.PathPrefix("/v1/").Subrouter()
	handler := h.handler.(*HTTPHandlerV1)
	sr.HandleFunc("/store/{req:[\\w-]*}", handler.SetStorePerformance).Methods("PUT")
	return nil
}
