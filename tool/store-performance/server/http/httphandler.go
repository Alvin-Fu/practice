package http

import (
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	myhttp "practice/lib/httpsvr"
	"practice/tool/store-performance/model"
	"practice/tool/store-performance/service"
	"strconv"
)

type HTTPHandlerV1 struct {
	storeService *service.StoreService
}

func (h *HTTPHandlerV1) GetPerformance(w http.ResponseWriter, r *http.Request) {}

func (h *HTTPHandlerV1) SetStorePerformance(w http.ResponseWriter, r *http.Request) {
	fmt.Println("jhsadfsaldk")
	body, err := ioutil.ReadAll(io.LimitReader(r.Body, 2048))
	if err != nil {
		myhttp.Response(w, http.StatusBadRequest, fmt.Sprintf("read put data fail: %s", err), []byte{})
		return
	}
	if len(body) == 0 {
		myhttp.Response(w, http.StatusBadRequest, "invalid empty body", []byte{})
		return
	}
	performance := &model.Performance{}
	json.Unmarshal(body, performance)
	data, err := h.storeService.GetStorePerformance(performance)
	if err != nil {
		myhttp.Response(w, http.StatusBadRequest, "fail", []byte(err.Error()))
		return
	}
	myhttp.Response(w, http.StatusOK, "ok", data)
	return
}

func (h *HTTPHandlerV1) SetStore(w http.ResponseWriter, r *http.Request) {
	htmlStr := "<html><head><title>业绩</title></head>" +
		"<body><form action='/v1/store' method=\"post\" enctype=\"multipart/form-data\">" +
		"<label>输入本月业绩</label>" + ":" +
		"<input type=\"text\" name='CurrentTurnover'  /> " +
		"<label>输入本月目标</label>" + ":" +
		"<input type=\"text\" name='TargetTurnover'  /> " +
		"<label>输入上月业绩</label>" + ":" +
		"<input type=\"text\" name='LastTurnover'  /> " +
		"<label>输入去年本月业绩</label>" + ":" +
		"<input type=\"text\" name='LastYearCurrentTurnover'  /> " +
		"<label>输入本月单数</label>" + ":" +
		"<input type=\"text\" name='Number'  /> " +
		"<label><input type=\"submit\" value=\"提交\"/></label></form></body></html>"
	if r.Method == "GET" {
		io.WriteString(w, htmlStr)
		return
	}

	number, _ := strconv.Atoi(r.FormValue("Number"))
	currentTurnover, _ := strconv.ParseFloat(r.FormValue("CurrentTurnover"), 64)
	targetTurnover, _ := strconv.ParseFloat(r.FormValue("TargetTurnover"), 64)
	lastTurnover, _ := strconv.ParseFloat(r.FormValue("LastTurnover"), 64)
	lastYearCurrentTurnover, _ := strconv.ParseFloat(r.FormValue("LastYearCurrentTurnover"), 64)

	pr := &model.Performance{
		Number:                  int64(number),
		CurrentTurnover:         currentTurnover,
		TargetTurnover:          targetTurnover,
		LastTurnover:            lastTurnover,
		LastYearCurrentTurnover: lastYearCurrentTurnover,
	}
	data, _ := h.storeService.GetStorePerformance(pr)
	io.WriteString(w, "计算成功: "+string(data))
}

func (h *HTTPHandlerV1) SetStoreHtml(w http.ResponseWriter, r *http.Request) {
	htmlData, _ := ioutil.ReadFile("../static/store.html")
	if r.Method == "GET" {
		io.WriteString(w, string(htmlData))
		return
	}
	number, _ := strconv.Atoi(r.FormValue("Number"))
	currentTurnover, _ := strconv.ParseFloat(r.FormValue("CurrentTurnover"), 64)
	targetTurnover, _ := strconv.ParseFloat(r.FormValue("TargetTurnover"), 64)
	lastTurnover, _ := strconv.ParseFloat(r.FormValue("LastTurnover"), 64)
	lastYearCurrentTurnover, _ := strconv.ParseFloat(r.FormValue("LastYearCurrentTurnover"), 64)

	pr := &model.Performance{
		Number:                  int64(number),
		CurrentTurnover:         currentTurnover,
		TargetTurnover:          targetTurnover,
		LastTurnover:            lastTurnover,
		LastYearCurrentTurnover: lastYearCurrentTurnover,
	}
	data, _ := h.storeService.GetStorePerformance(pr)
	io.WriteString(w, "计算成功: "+string(data))
}
