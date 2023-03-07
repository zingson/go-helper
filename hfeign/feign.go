package hfeign

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"github.com/sirupsen/logrus"
	"io/ioutil"
	"net/http"
	"net/url"
	"os"
	"strings"
	"time"
)

var (
	H_HTTP_PROXY = os.Getenv("H_HTTP_PROXY")
)

// Client 根据环境变量判断是否使用HTTP代理
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

func Request(ctx context.Context, method, url string, body string) (bytes []byte, err error) {
	var ms = time.Now().UnixMilli()
	defer func() {
		logrus.WithField("ms", fmt.Sprintf("%d", time.Now().UnixMilli()-ms)).Infof("FeignRequest %s  url: %s   req: %s  res: %s ", method, url, body, string(bytes))
	}()

	request, err := http.NewRequest(method, url, strings.NewReader(body))
	if err != nil {
		return nil, err
	}
	request.Header.Set("Content-Type", "application/json")

	var client = Client()

	if ctx != nil {
		if ctx.Value("token") != nil {
			request.Header.Set("token", ctx.Value("token").(string))
		}
		appid := ctx.Value("appid")
		if appid != nil && appid != "" {
			request.Header.Set("appid", appid.(string))
		}
		ctxClient := ctx.Value("client")
		if ctxClient != nil {
			client = ctxClient.(*http.Client)
		}
	}
	response, err := client.Do(request)
	if err != nil {
		return nil, err
	}

	bytes, err = ioutil.ReadAll(response.Body)
	if err != nil {
		return nil, err
	}
	if bytes != nil {
		var ferr *FeignErr
		err = json.Unmarshal(bytes, &ferr)
		if err != nil {
			return nil, err
		}
		if ferr.Errno == "00000" {
			return
		}
		return nil, errors.New(ferr.Errno + ":" + ferr.Error)
	}
	return nil, errors.New("99999:HTTP异常(" + response.Status + ")，请检查业务服务 " + method + " " + url)
}

type FeignErr struct {
	Errno string `json:"errno"`
	Error string `json:"error"`
}

func Do(ctx context.Context, method string, body string) ([]byte, error) {
	if method == "" {
		return nil, errors.New("接口名称不能为空")
	}
	name := strings.Split(method, ".")[1]
	if method[0:1] != "/" {
		method = "/" + method
	}
	reqUrl := "http://" + name + method
	return Request(ctx, "POST", reqUrl, body)
}

func Call(ctx context.Context, method string, i interface{}, o interface{}) (err error) {
	bytes, err := json.Marshal(i)
	if err != nil {
		return err
	}
	resp, err := Do(ctx, method, string(bytes))
	if err != nil {
		return err
	}
	err = json.Unmarshal(resp, o)
	return err
}

// InternalRequest 内部接口请求
func InternalRequest(ctx context.Context, method string, path string, i interface{}, o interface{}) (err error) {
	bytes, err := json.Marshal(i)
	if err != nil {
		return err
	}
	resp, err := Request(ctx, method, "http://"+path, string(bytes))
	if err != nil {
		return err
	}
	err = json.Unmarshal(resp, o)
	return err
}
