package metro2

import (
	"crypto/md5"
	"crypto/rand"
	"crypto/sha1"
	"encoding/json"
	"errors"
	"fmt"
	"github.com/sirupsen/logrus"
	"io"
	"io/ioutil"
	"net/http"
	"strings"
	"time"
)

// TokenKey Token
func TokenKey(token string) string {
	token = fmt.Sprintf("%x", sha1.Sum([]byte(token)))
	return "metro2:" + token
}

//HttpPost 调用接口
func HttpPost(config *Config, rid, path, reqData string) (resData string, err error) {
	var (
		beg        = time.Now().UnixMilli()
		serviceUrl = config.ServiceUrl + path
		reqHeader  = make(map[string]string)
		resHeader  = make(map[string]string)
	)
	defer func() {
		if e := recover(); e != nil {
			err = e.(error)
		}

		errMsg := ""
		if err != nil {
			errMsg = "ResponseError：" + err.Error()
		}
		logrus.Infof("rid: %s Metro2 HttpPost \nServiceUrl：%s  \nRequestHeader：%s  \nRequestBody: %s \nResponseHeader：%s  \nResponseBody: %s  \n%s  \n|%sms\n", rid, serviceUrl, jsonStringify(reqHeader), reqData, jsonStringify(resHeader), resData, errMsg, fmt.Sprintf("%d", time.Now().UnixMilli()-beg))
	}()

	req, err := http.NewRequest(http.MethodPost, serviceUrl, strings.NewReader(reqData))
	if err != nil {
		return
	}
	var (
		appid     = config.Appid
		nonce     = rid
		timestamp = fmt.Sprintf("%d", time.Now().UnixMilli())
		signtype  = "RSA2"
		version   = "2.0"
	)
	req.Header["appid"] = []string{config.Appid}
	req.Header["nonce"] = []string{nonce}
	req.Header["timestamp"] = []string{timestamp}
	req.Header["signtype"] = []string{signtype}
	req.Header["version"] = []string{version}
	req.Header["signature"] = []string{rsaSign(config.RsaPri, fmt.Sprintf("appid=%s&body=%s&nonce=%s&timestamp=%s&version=%s", appid, reqData, nonce, timestamp, version))}
	for k, v := range req.Header {
		reqHeader[k] = v[0]
	}
	req.Header.Set("Content-Type", "application/json")

	res, err := http.DefaultClient.Do(req)
	if err != nil {
		return
	}
	if res.StatusCode != http.StatusOK {
		err = errors.New(fmt.Sprintf("MTR_HTTP_%d:%s", res.StatusCode, res.Status))
		return
	}
	resData = string(must[[]byte](io.ReadAll(res.Body)))

	// 响应报文 异常时，不验证签名，直接返回错误
	if !strings.Contains(resData, "\"Code\":200,") {
		var rsd *ResBody
		rsd, err = jsonParse[*ResBody](resData)
		if err != nil {
			return
		}
		err = errors.New(fmt.Sprintf("MTR_%v:%s", rsd.Code, rsd.Message))
		return
	}

	for k, v := range res.Header {
		resHeader[k] = v[0]
	}
	resHeaderSign := res.Header.Get("signature")

	if err = rsaVerify(config.MetroRsaPub, resHeaderSign, resData); err != nil {
		err = errors.New("MTR_ERR:签名验证失败 " + err.Error())
		return
	}
	return
}

type ResBody struct {
	GuidRequest string `json:"GuidRequest"`
	Code        int64  `json:"Code"`
	Message     string `json:"Message"`
	Data        any    `json:"data"`
}

// Rand32 使用crypto/rand 随机赋值byte数组， 然后md5返回32位十六进制字符串
func Rand32() string {
	b := make([]byte, 48)
	if _, err := rand.Read(b); err != nil {
		return ""
	}
	return fmt.Sprintf("%x", md5.Sum(b))
}

//rsaSign 签名
func rsaSign(pri string, req string) (sign string) {
	logrus.Debugf("metro2 请求报文签名字符串：%s", req)
	sign, err := RsaSign(req, pri)
	if err != nil {
		panic(err)
	}
	return
}

//rsaVerify 验签
func rsaVerify(pub string, sign, res string) (err error) {
	logrus.Debugf("metro2 响应报文验签字符串：%s", res)
	return RsaVerify(sign, res, pub)
}

func jsonStringify(v any) string {
	b, _ := json.Marshal(v)
	return string(b)
}

// Parse 字符串解析为JSON对象
func jsonParse[T any](s string) (r T, err error) {
	err = json.Unmarshal([]byte(s), &r)
	if err != nil {
		return
	}
	return
}

func must[T any](v T, err error) T {
	if err != nil {
		panic(err)
	}
	return v
}

//httpCallback 接收地铁通知
func httpCallback(config *Config, rid string, request *http.Request, response http.ResponseWriter, f func(reqBody string) (resData any, err error)) {
	var (
		beg       = time.Now().UnixMilli()
		err       error
		sign      = request.Header.Get("signature")
		nonce     = request.Header.Get("nonce")
		timestamp = request.Header.Get("timestamp")
		version   = request.Header.Get("version")
		reqBody   = string(must[[]byte](ioutil.ReadAll(request.Body)))
		signStr   = fmt.Sprintf("appid=%s&body=%s&nonce=%s&timestamp=%s&version=%s", config.Appid, reqBody, nonce, timestamp, version)

		resBody string
	)

	defer func() {
		if e := recover(); e != nil {
			err = e.(error)
		}

		errMsg := ""
		if err != nil {
			errMsg = "ResponseError: " + err.Error()
		}
		logrus.Infof("rid: %s metro2 httpCallback path:%s  RequestHead: %s   ResponseHead: %s  %s", rid, request.RequestURI, jsonStringify(request.Header), jsonStringify(response.Header()))
		logrus.Infof("rid: %s metro2 httpCallback path:%s  RequestBody: %s   ResponseBody: %s  %s  | %sms", rid, request.RequestURI, reqBody, resBody, errMsg, fmt.Sprintf("%d", time.Now().UnixMilli()-beg))

		if err != nil {
			_, _ = response.Write([]byte(jsonStringify(&ResBody{
				GuidRequest: rid,
				Code:        500,
				Message:     err.Error(),
			})))
			return
		}

	}()

	// 验证签名
	if err = rsaVerify(config.MetroRsaPub, sign, signStr); err != nil {
		err = errors.New("METRO2:验签失败")
		return
	}

	// 调用业务参数
	resData, err := f(reqBody)
	if err != nil {
		return
	}

	resBody = jsonStringify(&ResBody{
		GuidRequest: rid,
		Code:        200,
		Message:     "ok",
		Data:        resData,
	})

	response.Header().Set("appid", config.Appid)
	response.Header().Set("timestamp", timestamp)
	response.Header().Set("nonce", rid)
	response.Header().Set("signature", rsaSign(config.RsaPri, resBody))
	_, _ = response.Write([]byte(resBody))
}
