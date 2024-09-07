package api

import (
	"bytes"
	"fmt"
	"io"
	"net/http"
	"sort"

	"github.com/lcsin/doraemon/util"
)

// POST 发送HTTP POST请求
func POST(url string, body []byte) ([]byte, error) {
	buf := bytes.NewBuffer(body)
	resp, err := http.Post(url, "application/json;charset=utf-8", buf)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	respBody, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("%v", string(respBody))
	}

	return respBody, nil
}

// GET 发送HTTP GET请求
func GET(url string) ([]byte, error) {
	resp, err := http.Get(url)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	respBody, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("%v", string(respBody))
	}

	return respBody, nil
}

// Sign 接口签名
func Sign(params map[string]string, key string) string {
	var fields []string
	for k := range params {
		fields = append(fields, k)
	}
	sort.Strings(fields)

	var sign string
	for _, v := range fields {
		sign += v
		val, _ := params[v]
		sign += val
	}
	sign += key

	return util.MD5(sign)
}
