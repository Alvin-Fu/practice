package testsvr

import (
	"encoding/json"
	"fmt"
	"github.com/gorilla/mux"
	"net/http"
)

type HTTPHandlerV1 struct{}

func (h *HTTPHandlerV1) GetMatchInfo(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Get all match info!")
	matchId := mux.Vars(r)["mid"]
	fmt.Printf("Match id: %s\n", matchId)
	h.response(w, http.StatusOK, "ok", []byte("All match info!"))
}

func (h *HTTPHandlerV1) GetOneMatchInfo(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Get one match info!")
	matchId := mux.Vars(r)["mid"]
	fmt.Printf("Match id: %s\n", matchId)
	h.response(w, http.StatusOK, "ok", []byte("One match info!"))
}

// response make response in v1 and send to client.
func (h *HTTPHandlerV1) response(w http.ResponseWriter, statusCode int, errmsg string, respJSON []byte) error {
	// use statusCode as error code
	w.Header().Add("Content-Type", "application/json; charset=utf-8")
	w.WriteHeader(statusCode)

	resp, _ := json.Marshal(struct {
		Errmsg string
		Data   string
	}{errmsg, string(respJSON)})

	_, err := w.Write(resp)
	return err
}
