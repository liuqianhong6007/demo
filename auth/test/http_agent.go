package test

import (
	"bytes"
	"net/http"
	"net/url"
	"strings"
	"time"
)

func NewHttpAgent() *HttpAgent {
	return &HttpAgent{
		timeout: 2 * time.Second,
	}
}

// http 代理
type HttpAgent struct {
	timeout  time.Duration         // 超时时长
	signFunc func(r *http.Request) // 签名函数
}

// 设置超时时长
func (h *HttpAgent) SetTimeout(timeout time.Duration) {
	h.timeout = timeout
}

// 设置签名函数
func (h *HttpAgent) SetSignFunc(signFunc func(r *http.Request)) {
	h.signFunc = signFunc
}

// 发送 get 请求
func (h *HttpAgent) Get(path string, header http.Header, params map[string]string) (response *http.Response, err error) {
	if !strings.HasSuffix(path, "?") {
		path = path + "?"
	}

	if params != nil {
		for key, value := range params {
			path = path + key + "=" + value + "&"
		}
		if strings.HasSuffix(path, "&") {
			path = path[:len(path)-1]
		}
	}

	client := http.Client{}
	client.Timeout = h.timeout
	request, errNew := http.NewRequest(http.MethodGet, path, nil)
	if errNew != nil {
		err = errNew
		return
	}
	request.Header = header
	if h.signFunc != nil {
		h.signFunc(request)
	}

	resp, errDo := client.Do(request)
	if errDo != nil {
		err = errDo
	} else {
		response = resp
	}
	return
}

// 发送 post 请求
func (h *HttpAgent) Post(path string, header http.Header, body []byte) (response *http.Response, err error) {
	client := http.Client{}
	client.Timeout = h.timeout
	request, errNew := http.NewRequest(http.MethodPost, path, bytes.NewReader(body))
	if errNew != nil {
		err = errNew
		return
	}
	request.Header = header
	if h.signFunc != nil {
		h.signFunc(request)
	}

	resp, errDo := client.Do(request)
	if errDo != nil {
		err = errDo
	} else {
		response = resp
	}
	return
}

// 发送 post form 请求
func (h *HttpAgent) PostForm(path string, header http.Header, value url.Values) (response *http.Response, err error) {
	client := http.Client{}
	client.Timeout = h.timeout
	request, errNew := http.NewRequest(http.MethodPost, path, strings.NewReader(value.Encode()))
	if errNew != nil {
		err = errNew
		return
	}
	if header != nil {
		request.Header = header
	}
	request.Header.Set("Content-Type", "application/x-www-form-urlencoded")
	if h.signFunc != nil {
		h.signFunc(request)
	}

	resp, errDo := client.Do(request)
	if errDo != nil {
		err = errDo
	} else {
		response = resp
	}
	return
}
