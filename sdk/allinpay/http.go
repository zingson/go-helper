package allinpay

import (
	"crypto/md5"
	"crypto/rand"
	"encoding/json"
	"errors"
	"fmt"
	"github.com/sirupsen/logrus"
	"io/ioutil"
	"net/http"
	"net/url"
	"strconv"
	"time"
)

// PostForm 表单POST请求
func PostForm(conf *Config, path string, bm map[string]string, result interface{}) (err error) {
	if conf == nil {
		err = errors.New("收银宝接口配置不能为空")
		return
	}
	var (
		bmilli  = time.Now().UnixMilli()
		rid     = Rand32()
		reqUrl  = conf.BaseUrl + path
		reqJson string
		resJson string
	)
	defer func() {
		var errMsg string
		if err != nil {
			errMsg = "异常：" + err.Error() + " ,rid=" + rid
		}
		millisecond := fmt.Sprintf("%d", time.Now().UnixMilli()-bmilli)
		logrus.WithField("rid", rid).WithField("millisecond", millisecond).Infof("sdk allinpay 收银宝 请求URL：%s  请求报文：%s  响应报文：%s  %s  | %sms", reqUrl, reqJson, resJson, errMsg, millisecond)
	}()

	if bm == nil {
		bm = make(map[string]string)
	}

	bm["cusid"] = conf.Cusid
	bm["appid"] = conf.Appid
	bm["version"] = VERSION_11
	bm["randomstr"] = Rand32()
	bm["signtype"] = SIGN_TYPE_RSA

	sign, err := RsaSign(conf, bm)
	if err != nil {
		return
	}
	bm["sign"] = sign

	reqJsonBytes, err := json.Marshal(bm)
	reqJson = string(reqJsonBytes)

	data := url.Values{}
	for k, v := range bm {
		if v == "" {
			continue
		}
		data.Set(k, v)
	}
	resp, err := http.PostForm(reqUrl, data)
	if err != nil {
		return
	}
	ret, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return
	}
	resJson = string(ret)

	var retBm = make(map[string]string)
	err = json.Unmarshal(ret, &retBm)
	if err != nil {
		return
	}

	// 非成功状态，不验证签名
	if retBm["retcode"] == string(RET_FAIl) {
		err = errors.New(retBm["retcode"] + ":" + retBm["retmsg"])
		return
	}

	err = RsaVerify(conf, retBm)
	if err != nil {
		return
	}

	err = json.Unmarshal(ret, &result)
	if err != nil {
		return
	}
	return
}

// Rand32 使用crypto/rand 随机赋值byte数组， 然后md5返回32位十六进制字符串
func Rand32() string {
	b := make([]byte, 48)
	if _, err := rand.Read(b); err != nil {
		return ""
	}
	return fmt.Sprintf("%x", md5.Sum(b))
}

//ParseAmt 解析金额
func ParseAmt(amt string) int64 {
	v, err := strconv.ParseInt(amt, 64, 10)
	if err != nil {
		v = 0
	}
	return v
}
