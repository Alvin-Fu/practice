package http

import (
	myhttp "domino/lib/httpsvr"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"practice/tool/store-performance/service"
)

type HTTPHandlerV1 struct {
	storeService *service.StoreService
}

func (h *HTTPHandlerV1) SetStorePerformance(w http.ResponseWriter, r *http.Request) {
	body, err := ioutil.ReadAll(io.LimitReader(r.Body, 2048))
	if err != nil {
		myhttp.Response(w, http.StatusBadRequest, fmt.Sprintf("read put data fail: %s", err), []byte{})
		return
	}
	if len(body) == 0 {
		myhttp.Response(w, http.StatusBadRequest, "invalid empty body", []byte{})
		return
	}
	data, err := h.storeService.GetStorePerformance(body)
	if err != nil {
		myhttp.Response(w, http.StatusBadRequest, "fail", []byte(err.Error()))
		return
	}
	myhttp.Response(w, http.StatusOK, "ok", data)
	return
}
