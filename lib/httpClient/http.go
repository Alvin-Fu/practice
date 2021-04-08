package httpClient

import (
	"bytes"
	"encoding/json"
	"io"
	"io/ioutil"
	"net"
	"net/http"
	"time"
)

const (
	MaxIdleConns        int           = 5000
	MaxIdleConnsPerHost int           = 100
	IdleConnTimeout     int           = 90
	ConnTimeout         time.Duration = 10
)

var (
	httpClient    *http.Client
	contentType   *ContentType
	requestMethod *RequestMethod
)

type (
	HttpResult struct {
		StatusCode   int
		ResponseBody []byte
	}

	ContentType struct {
		FormData              string
		XWwwFormUrlEncode     string
		TextPlain             string
		ApplicationJson       string
		ApplicationJavaScript string
		ApplicationXml        string
		TextXml               string
		TextHtml              string
	}

	RequestMethod struct {
		GET    string
		POST   string
		PUT    string
		PATCH  string
		DELETE string
	}
)

func init() {
	httpClient = createHTTPClient()
	contentType = &ContentType{
		FormData:              "multipart/form-data",
		XWwwFormUrlEncode:     "application/x-www-form-urlencoded",
		TextHtml:              "text/html",
		ApplicationJavaScript: "application/javascript",
		ApplicationJson:       "application/json",
		ApplicationXml:        "application/xml",
		TextPlain:             "text/plain",
		TextXml:               "text/xml",
	}
	requestMethod = &RequestMethod{
		GET:    "GET",
		POST:   "POST",
		PUT:    "PUT",
		PATCH:  "PATCH",
		DELETE: "DELETE",
	}
}

// createHTTPClient for connection re-use
func createHTTPClient() *http.Client {

	client := &http.Client{
		Transport: &http.Transport{
			//DisableKeepAlives:false,// 是否开启http keepalive功能，也即是否重用连接，默认开启(false)
			Proxy: http.ProxyFromEnvironment,
			// 通过设置tls.Config的InsecureSkipVerify为true，client将不再对服务端的证书进行校验
			//TLSClientConfig:    &tls.Config{InsecureSkipVerify: true},
			DialContext: (&net.Dialer{
				Timeout:   30 * time.Second,
				KeepAlive: 30 * time.Second,
			}).DialContext, // 该函数用于创建http（非https）连接，通常需要关注Timeout和KeepAlive参数
			MaxIdleConns:        MaxIdleConns,        // 连接池对所有host的最大链接数量，host也即dest-ip，默认为无穷大（0）
			MaxIdleConnsPerHost: MaxIdleConnsPerHost, // 连接池对每个host的最大链接数量
			// 空闲timeout设置，也即socket在该时间内没有交互则自动关闭连接
			// （注意：该timeout起点是从每次空闲开始计时，若有交互则重置为0）,
			// 该参数通常设置为分钟级别，例如：90秒
			IdleConnTimeout: time.Duration(IdleConnTimeout) * time.Second,
		},
		Timeout: ConnTimeout * time.Second,
	}
	return client
}

func Get(url string) (result HttpResult, err error) {
	return httpRequest(requestMethod.GET, url, bytes.NewBufferString(""), "", nil)
}

func PostJson(url string, data interface{}, header interface{}) (result HttpResult, err error) {
	bodyStr, _ := json.Marshal(data)
	ret := bytes.NewBuffer(bodyStr)
	return post(url, ret, contentType.ApplicationJson, header)
}

func PostForm(url string, data io.Reader, header interface{}) (result HttpResult, err error) {
	return post(url, data, contentType.XWwwFormUrlEncode, header)
}

func post(url string, body io.Reader, contentType string, header interface{}) (result HttpResult, err error) {
	return httpRequest(requestMethod.POST, url, body, contentType, header)
}

func httpRequest(method string, url string, body io.Reader, contentType string, header interface{}) (result HttpResult, err error) {
	req, err := http.NewRequest(method, url, body)
	if err != nil {
		return
	}
	if contentType != "" {
		req.Header.Add("content-type", contentType)
	}
	if header != nil {
		headerStr, _ := json.Marshal(header)
		var mapResult map[string]string
		err := json.Unmarshal([]byte(headerStr), &mapResult)
		if err != nil {
			panic(err)
		}
		for key, val := range mapResult {
			req.Header.Add(key, val)
		}
	}
	defer req.Body.Close()
	resp, err := httpClient.Do(req)
	if err != nil {
		return
	}
	defer resp.Body.Close()
	res, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return
	}
	result.StatusCode = resp.StatusCode
	result.ResponseBody = res
	return
}
