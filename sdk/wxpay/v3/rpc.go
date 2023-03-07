package v3

import (
	"bytes"
	"crypto/rand"
	"encoding/json"
	"errors"
	"fmt"
	"github.com/sirupsen/logrus"
	"io"
	"math/big"
	"net/http"
	"strconv"
	"strings"
	"time"
)

// Authorization 拼接权限验证字符串
func Authorization(c *Client, method, path, body string) (authorization string, err error) {
	authType := "WECHATPAY2-SHA256-RSA2048" //固定字符串
	timestamp := strconv.FormatInt(time.Now().Unix(), 10)
	nonceStr := RandomString(32)
	signature, err := V3Sign(method, path, body, timestamp, nonceStr, c.mchPriKey)
	if err != nil {
		return
	}
	authorization = fmt.Sprintf(`%s mchid="%s",nonce_str="%s",signature="%s",timestamp="%s",serial_no="%s"`, authType, c.Mchid, nonceStr, signature, timestamp, c.SerialNo)
	return
}

//Call 需要验签的接口使用此方法调用
func Call(c *Client, method, path string, i interface{}, o interface{}) (err error) {
	var (
		reqBody string
		resBody string
	)
	if i != nil {
		reqBytes, _ := json.Marshal(i)
		reqBody = string(reqBytes)
	}

	rBytes, header, err := Do(c, method, path, reqBody)
	if err != nil {
		return
	}
	resBody = string(rBytes)

	//requestId := resp.Header.Get("Request-ID")
	signature := header.Get("Wechatpay-Signature")
	serial := header.Get("Wechatpay-Serial")
	timestamp := header.Get("Wechatpay-Timestamp")
	nonce := header.Get("Wechatpay-Nonce")

	pubKey, err := c.GetWxPubKey(serial)
	if err != nil {
		return
	}
	ok, err := V3SignVery(signature, timestamp, nonce, resBody, pubKey)
	if err != nil {
		return
	}
	if !ok {
		return errors.New("签名校验失败")
	}
	err = json.Unmarshal(rBytes, o)
	if err != nil {
		return
	}
	return
}

// Do 发送接口请求，无需验证签名的可直接调用
func Do(c *Client, method, path, reqBody string) (rBytes []byte, header http.Header, err error) {
	var (
		milli   = time.Now().UnixMilli()
		reqUrl  = c.ServiceUrl + path
		resBody string
		rid     = RandomString(32)
	)
	defer func() {
		errMsg := ""
		if err != nil {
			errMsg = "响应异常：" + err.Error()
		}
		logrus.
			WithField("rid", rid).
			WithField("ms", fmt.Sprintf("%d", time.Now().UnixMilli()-milli)).
			Infof("wxpay 请求URL：%s %s  请求报文：%s  响应报文：%s  %s ", method, reqUrl, reqBody, resBody, errMsg)
	}()

	authorization, err := Authorization(c, method, path, reqBody)
	if err != nil {
		return
	}

	request, err := http.NewRequest(method, reqUrl, strings.NewReader(reqBody))
	if err != nil {
		return
	}
	request.Header.Set("User-Agent", "v3;helper")
	request.Header.Set("Content-Type", "application/json")
	request.Header.Set("Accept", "application/json")
	request.Header.Set("Authorization", authorization)
	request.Header.Set("Wechatpay-Serial", c.wxSerialNo) // 存在敏感字段加密时必填，其它场景可选
	resp, err := http.DefaultClient.Do(request)
	if err != nil {
		return
	}
	header = resp.Header
	rid = header.Get("Request-ID")

	// HTTP 返回204，处理成功，应答无内容
	if resp.StatusCode == 204 {
		return
	}

	rBytes, err = io.ReadAll(resp.Body)
	if err != nil {
		return
	}
	resBody = string(rBytes)

	// HTTP 返回非200，直接返回错误
	if resp.StatusCode != 200 {
		var eres *ErrResponse
		err = json.Unmarshal(rBytes, &eres)
		if err != nil {
			return
		}
		err = errors.New(eres.Code + ":" + eres.Message)
		return
	}
	return
}

type ErrResponse struct {
	Code    string `json:"code"`
	Message string `json:"message"`
}

// RandomString 生成随机字符串
func RandomString(len int) string {
	var container string
	var str = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ1234567890"
	b := bytes.NewBufferString(str)
	length := b.Len()
	bigInt := big.NewInt(int64(length))
	for i := 0; i < len; i++ {
		randomInt, _ := rand.Int(rand.Reader, bigInt)
		container += string(str[randomInt.Int64()])
	}
	return container
}
