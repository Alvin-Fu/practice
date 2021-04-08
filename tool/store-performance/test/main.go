package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"practice/tool/store-performance/model"
	"strings"
)

func main() {
	performance := &model.Performance{
		TargetTurnover:          390000,
		CurrentTurnover:         318444.4,
		LastTurnover:            311730,
		LastYearCurrentTurnover: 402686.3,
		Number:                  372,
	}
	data, _ := json.Marshal(performance)
	Put(data)
}

func Put(reqData []byte) {
	url := "http://192.168.28.176:8090/v1/store/req"
	payload := strings.NewReader(string(reqData))
	req, _ := http.NewRequest("PUT", url, payload)
	req.Header.Add("Authorization", "token_value")
	req.Header.Add("Content-Type", "text/plain; charset=UTF-8")
	req.Header.Add("User-Agent", "Go-http-client/1.14")
	req.Header.Add("Transfer-Encoding", "chunked")
	req.Header.Add("Accept-Encoding", "gzip, deflate")
	res, err := http.DefaultClient.Do(req)
	if err != nil {
		fmt.Println(err)
		return
	}
	defer res.Body.Close()
	body, _ := ioutil.ReadAll(res.Body)
	fmt.Println(string(body))
}
