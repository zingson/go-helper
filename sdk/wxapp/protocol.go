package wxapp

import (
	"bytes"
	"encoding/json"
	"errors"
	"github.com/sirupsen/logrus"
	"io"
	"net/http"
	"strconv"
)

// 微信小程序，接口GET请求
func get[T any](url string) (res T, err error) {
	var (
		resBody string
		errMsg  string
	)
	defer func() {
		if e := recover(); e != nil {
			err = e.(error)
		}
		if err != nil {
			errMsg = "\n错误信息：" + err.Error() + "\n"
		}
		logrus.Infof("wxapp GET \n请求地址：%s \n响应报文：%s  %s", url, resBody, errMsg)
	}()

	resp, err := http.Get(url)
	if err != nil {
		return
	}
	b, err := io.ReadAll(resp.Body)
	if err != nil {
		return
	}
	resBody = string(b)

	var rs Result
	err = json.Unmarshal(b, &rs)
	if err != nil {
		return
	}
	if rs.Errcode != 0 {
		err = errors.New(strconv.FormatInt(rs.Errcode, 10) + ":" + rs.Errmsg)
		return
	}

	err = json.Unmarshal(b, &res)
	if err != nil {
		return
	}
	return
}

// 微信小程序，接口POST请求
func post[T any](url string, req any) (res T, err error) {
	var (
		reqBody string
		resBody string
		errMsg  string
	)
	defer func() {
		if e := recover(); e != nil {
			err = e.(error)
		}
		if err != nil {
			errMsg = "\n错误信息：" + err.Error() + "\n"
		}
		logrus.Infof("wxapp POST \n请求地址：%s \n请求报文：%s \n响应报文：%s  %s", url, reqBody, resBody, errMsg)
	}()

	b, _ := json.Marshal(req)
	reqBody = string(b)
	resp, err := http.Post(url, "application/json", bytes.NewReader(b))
	if err != nil {
		return
	}
	respBytes, err := io.ReadAll(resp.Body)
	if err != nil {
		return
	}
	resBody = string(respBytes)

	var rs Result
	err = json.Unmarshal(respBytes, &rs)
	if err != nil {
		return
	}
	if rs.Errcode != 0 {
		err = errors.New(strconv.FormatInt(rs.Errcode, 10) + ":" + rs.Errmsg)
		return
	}

	err = json.Unmarshal(respBytes, &res)
	if err != nil {
		return
	}
	return
}

type Result struct {
	Errcode int64  `json:"errcode"` // 0=成功，其它失败,成功时没有这个字段，通过默认的0值判断成功
	Errmsg  string `json:"errmsg"`
}
