package bytedancepay

import (
	"encoding/json"
	"errors"
	"github.com/sirupsen/logrus"
	"io/ioutil"
	"net/http"
	"strings"
	"time"
)

// 接口请求
func post(conf *Config, path string, reqBody string) (resBody string, err error) {
	var (
		url      = conf.ServiceUrl + path
		begMilli = time.Now().UnixMilli()
	)
	defer func() {
		errMsg := ""
		if err != nil {
			errMsg = "ZJERR:" + err.Error()
		}
		if e := recover(); e != nil {
			errMsg = "ZJERR:" + (e.(error)).Error()
		}
		logrus.
			WithField("mchid", conf.Mchid).
			WithField("appid", conf.AppId).
			Infof("字节小程序支付 POST %s  请求报文：%s  响应报文：%s  %s  %dms", url, reqBody, resBody, errMsg, time.Now().UnixMilli()-begMilli)
	}()
	resp, err := http.Post(url, "application/json", strings.NewReader(reqBody))
	if err != nil {
		return
	}
	if resp.StatusCode != http.StatusOK {
		err = errors.New(resp.Status)
		return
	}
	rb, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return
	}
	resBody = string(rb)
	return
}

func structToMap(in interface{}) (pmap map[string]interface{}) {
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
