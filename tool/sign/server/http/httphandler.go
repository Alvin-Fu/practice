package http

import (
	"net/http"
	myhttp "practice/lib/httpsvr"
	"practice/tool/sign/service"
)

type HTTPHandlerV1 struct {
	storeService *service.StoreService
}

func (h *HTTPHandlerV1) GetStudents(w http.ResponseWriter, r *http.Request) {

}

func (h *HTTPHandlerV1) GetStudentName(w http.ResponseWriter, r *http.Request) {
	myhttp.Response(w, http.StatusOK, "ok", []byte(h.storeService.GetOneStudentName()))
}
