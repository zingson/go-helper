package hhttp

import (
	"net/http"
	"net/url"
	"os"
	"time"
)

var (
	H_HTTP_PROXY = os.Getenv("H_HTTP_PROXY")
)

// Client Http Client，根据环境变量判断是否使用HTTP代理
func Client() *http.Client {
	if H_HTTP_PROXY == "" {
		return http.DefaultClient
	}
	return &http.Client{
		Transport: &http.Transport{
			Proxy: func(request *http.Request) (url *url.URL, err error) {
				request.URL.Host = H_HTTP_PROXY
				return request.URL, err
			},
		},
		Timeout: 5 * time.Second,
	}
}
