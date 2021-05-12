package svrmonitoring

import (
    "net/http"
    "fmt"
    "github.com/gorilla/mux"
    "io/ioutil"
    "io"
    "strconv"
    "encoding/json"
)

// HTTPHandlerV1 define all handler for request.
type HTTPHandlerV1 struct {
}

func (h *HTTPHandlerV1) Stat(w http.ResponseWriter, r *http.Request) {
    fmt.Fprintf(w, "home page of v1\n")
}

func (h *HTTPHandlerV1) GetRobots(w http.ResponseWriter, r *http.Request){}

func (h *HTTPHandlerV1) GetConnectRobots(w http.ResponseWriter, r *http.Request) {
    opt := mux.Vars(r)["opt"]
    body, err := ioutil.ReadAll(io.LimitReader(r.Body, 2048))
    if err != nil {
        h.response(w, http.StatusBadRequest, fmt.Sprintf("read put data fail: %s", err), []byte{})
        return
    }
    if len(body) == 0 {
        h.response(w, http.StatusBadRequest, "invalid empty body", []byte{})
        return
    }
    stat := make([]byte, 0)
    switch opt {
    case "hour":
        body := string(body)
        _, err := strconv.Atoi(body)
        if err != nil {
            h.response(w, http.StatusBadRequest, "invalid empty body", []byte("invalid empty body"))
            return
        }
    default:
        h.response(w, http.StatusBadRequest, fmt.Sprintf("invalid option: %s", opt), []byte("invalid option"))
        return
    }
    fmt.Println(string(stat))
    err = h.response(w, http.StatusOK, "ok", stat)
    if err != nil {
    }
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
