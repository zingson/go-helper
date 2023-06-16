package bank_zhywt

import (
	"encoding/json"
	"errors"
	"github.com/sirupsen/logrus"
	"io"
	"net/http"
	"net/url"
	"time"
)

const (
	Version  = "1.0"
	Charset  = "UTF-8"
	SignType = "SHA-256"

	//接口响应码

	SUC0000 = "SUC0000" // 成功
)

type RequestBody[T any] struct {
	Version  string `json:"version" ` //接口版本号,固定为”1.0”
	Charset  string `json:"charset"`  //参数编码,固定为“UTF-8”
	Sign     string `json:"sign"`     //报文签名,使用商户支付密钥对reqData内的数据进行签名
	SignType string `json:"signType"` //签名算法,固定为”SHA-256”
	ReqData  T      `json:"reqData"`  //请求数据

}

type ResponseBody[T any] struct {
	Version  string `json:"version" `
	Charset  string `json:"charset"`
	Sign     string `json:"sign"`
	SignType string `json:"signType"`
	RspData  T      `json:"rspData"` //响应数据
}

// Post 统一接口请求
func Post[P, R any](conf *Config, apiUrl string, reqData P) (rspData R, err error) {
	var (
		begMilli = time.Now().UnixMilli()
		reqBody  string
		resBody  string
	)
	defer func() {
		if e := recover(); e != nil {
			err = e.(error)
		}
		errMsg := ""
		if err != nil {
			errMsg = "\n错误信息：" + err.Error()
		}
		logrus.WithField("mchid", conf.BranchNo+conf.MerchantNo).Infof("招行一网通H5支付 \n接口：%s  \n请求：%s  \n响应：%s  %s  \n耗时：%dms", apiUrl, reqBody, resBody, errMsg, time.Now().UnixMilli()-begMilli)
	}()

	reqBody = jsonMarshal(&RequestBody[P]{
		Version:  Version,
		Charset:  Charset,
		Sign:     Sha256Sign(SortMap(StructToMap(reqData), true) + "&" + conf.Merkey),
		SignType: SignType,
		ReqData:  reqData,
	})

	values := url.Values{}
	values.Set("jsonRequestData", reqBody)
	values.Set("charset", "UTF-8")
	resp, err := http.PostForm(apiUrl, values)
	if err != nil {
		return
	}
	ret, err := io.ReadAll(resp.Body)
	if err != nil {
		return
	}
	resBody = string(ret)

	var responseBody *ResponseBody[R]
	err = json.Unmarshal(ret, &responseBody)
	if err != nil {
		return
	}
	rspData = responseBody.RspData
	return
}

func jsonMarshal(v any) string {
	b, _ := json.Marshal(v)
	return string(b)
}

type NoticeBody[T any] struct {
	Version    string `json:"version" `   //接口版本号,固定为”1.0”
	Charset    string `json:"charset"`    //参数编码,固定为“UTF-8”
	Sign       string `json:"sign"`       //报文签名,使用招行私钥对noticeData内的数据进行签名；商户需使用招行公钥验签。
	SignType   string `json:"signType"`   //签名算法,固定为”RSA”
	NoticeData T      `json:"noticeData"` //通知数据
}

func ParseNotice[T any](conf *Config, path string, reqBody []byte) (noticeData T, err error) {

	defer func() {
		resBody := "成功"
		errMsg := ""
		if e := recover(); e != nil {
			err = e.(error)
		}
		if err != nil {
			resBody = "异常"
			errMsg = "\n异常：" + err.Error()
		}
		logrus.WithField("mchid", conf.BranchNo+conf.MerchantNo).Infof("招行一网通H5支付 通知 \n接口：%s  \n请求：%s  \n响应：%s  %s ", path, string(reqBody), resBody, errMsg)
	}()

	var noticeBody *NoticeBody[T]
	err = json.Unmarshal(reqBody, &noticeBody)
	if err != nil {
		return
	}

	// 一网通RSA公钥
	pubKey, err := conf.GetFbPubKey()
	if err != nil {
		return
	}

	// 验签
	err = RsaVerify(noticeBody.Sign, SortMap(StructToMap(noticeBody.NoticeData), true), pubKey)
	if err != nil {
		logrus.Error("ParseNotice RsaVerify " + err.Error())
		err = errors.New("签名验证失败")
		return
	}

	noticeData = noticeBody.NoticeData
	return
}
