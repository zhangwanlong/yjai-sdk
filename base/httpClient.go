package base

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
	"strings"
	"time"
)

type httpClient struct {
	Apikey    string `json:"apikey,omitempty"`
	Apisecret string `json:"apisecret,omitempty"`
	TimeOut   int64  `json:"timeOut,omitempty"` // 超时
}

func NewHttpClient(apikey string, apisecret string, timeOut int64) *httpClient {
	return &httpClient{
		Apikey:    apikey,
		Apisecret: apisecret,
		TimeOut:   timeOut,
	}
}

func (h httpClient) RequestPost(requerstUrl string, values url.Values) (responseStr string, err error) {
	// 请求时的时间戳秒数
	now := fmt.Sprintf("%d", time.Now().Unix())
	t, err := json.Marshal(h)
	fmt.Println("----")
	fmt.Println(string(t))
	fmt.Println(h.Apikey)
	//return
	values.Add("apikey", h.Apikey)
	values.Add("timestamp", now)
	// 得到签名
	sign := GetSign(values, h.Apisecret)

	headerMap := make(map[string]string)
	headerMap["sign"] = sign
	headerMap["Content-Type"] = "application/x-www-form-urlencoded"

	// 发起请求
	responseStr, err = h.PostWithHeader(requerstUrl, values, headerMap, h.TimeOut)
	return
	if err != nil {
		return "", err
	}

	return responseStr, nil
}

func (h httpClient) PostWithHeader(url string, data url.Values, headerMap map[string]string, timeout int64) (string, error) {
	client := &http.Client{Timeout: time.Duration(timeout) * time.Second}
	req, err := http.NewRequest("POST", url, strings.NewReader(data.Encode()))
	if err != nil {
		return "", err
	}

	for k, v := range headerMap {
		req.Header.Set(k, v)
	}

	resp, err := client.Do(req)
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return "", err
	}
	return string(body), nil
}
