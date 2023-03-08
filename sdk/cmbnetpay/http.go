package cmbnetpay

import (
	"encoding/json"
	"github.com/sirupsen/logrus"
	"io/ioutil"
	"net/http"
	"net/url"
	"time"
)

//PostForm 表单POST请求
func PostForm(conf *Config, path, reqBody string) (resBody string, err error) {
	var (
		apiUrl   = path
		begMilli = time.Now().UnixMilli()
	)
	defer func() {
		errMsg := ""
		if err != nil {
			errMsg = "CMB_NET_PAY:" + err.Error()
		}
		if e := recover(); e != nil {
			errMsg = "CMB_NET_PAY:" + (e.(error)).Error()
		}
		logrus.WithField("mchid", conf.MerchantNo).Infof("招行一网通接口 POST %s  请求报文：%s  响应报文：%s  %s  %dms", apiUrl, reqBody, resBody, errMsg, time.Now().UnixMilli()-begMilli)
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
	ret, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return
	}
	resBody = string(ret)
	return
}

func StructToMap(in interface{}) (pmap map[string]string) {
	b, err := json.Marshal(in)
	if err != nil {
		panic(err)
	}
	err = json.Unmarshal(b, &pmap)
	if err != nil {
		panic(err)
	}
	return
}
