package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"strings"
)

//参考：https://www.cnblogs.com/Goden/p/4658287.html
func main() {
	url := "http://192.168.81.15:15000/v1/config/log-level"

	payload := strings.NewReader("debug")

	req, _ := http.NewRequest("PUT", url, payload)
	req.Header.Add("Authorization", "token_value")
	req.Header.Add("Content-Type", "text/plain; charset=UTF-8")
	req.Header.Add("User-Agent", "Go-http-client/1.14")
	req.Header.Add("Transfer-Encoding", "chunked")
	req.Header.Add("Accept-Encoding", "gzip, deflate")
	//req.URL.Query().Add("log-level", "debug")
	//req.PostForm.Set("log-level", "debug")
	res, _ := http.DefaultClient.Do(req)

	defer res.Body.Close()
	body, _ := ioutil.ReadAll(res.Body)

	fmt.Println(res)
	fmt.Println(string(body))
}
