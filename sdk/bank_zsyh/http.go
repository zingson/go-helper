package bank_zsyh

import (
	"encoding/json"
	"github.com/sirupsen/logrus"
	"io"
	"net/http"
	"net/url"
	"time"
)

// 删除 2023-06-14

// PostForm 表单POST请求
func PostForm(conf *Config, apiUrl, reqBody string) (resBody string, err error) {
	var (
		begMilli = time.Now().UnixMilli()
	)
	defer func() {
		if e := recover(); e != nil {
			err = e.(error)
		}
		errMsg := ""
		if err != nil {
			errMsg = "\n错误信息：" + err.Error()
		}
		logrus.WithField("mchid", conf.MerchantNo).Infof("招行一网通接口 \nPOST %s  \n请求报文：%s  \n响应报文：%s  %s  \n%dms", apiUrl, reqBody, resBody, errMsg, time.Now().UnixMilli()-begMilli)
	}()
	bm := make(map[string]string)
	bm["jsonRequestData"] = reqBody
	reqBodyJson, err := json.Marshal(bm)
	if err != nil {
		return
	}
	reqBody = string(reqBodyJson)
	data := url.Values{}
	for k, v := range bm {
		if v == "" {
			continue
		}
		data.Set(k, v)
	}
	resp, err := http.PostForm(apiUrl, data)
	if err != nil {
		return
	}
	ret, err := io.ReadAll(resp.Body)
	if err != nil {
		return
	}
	resBody = string(ret)
	return
}
