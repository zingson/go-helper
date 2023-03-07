package upapi

import (
	"encoding/hex"
	"encoding/json"
	"errors"
	"fmt"
	"github.com/sirupsen/logrus"
	"io/ioutil"
	"math/rand"
	"net/http"
	"strings"
	"time"
)

//Post 云闪付接口 POST请求
func Post(c *Config, path string, bodyMap *BodyMap) (respBody *RespBody, err error) {
	var (
		rid      = Rand32()
		nlog     = logrus.WithField("path", path).WithField("appid", c.Appid)
		begMilli = time.Now().UnixMilli()
		url      = c.ServiceUrl + path
		reqBody  string
		resBody  string
		errMsg   string
	)
	defer func() {
		//fmt.Println("请求URL：" + url)
		//fmt.Println("请求报文：" + reqBody)
		//fmt.Println("响应报文：" + resBody)
		if err != nil {
			errMsg = "Error: " + err.Error()
		}
		ms := fmt.Sprintf("%d", time.Now().UnixMilli()-begMilli)
		nlog.WithField("ms", ms).Infof("云闪付upapi \n请求URL：POST %s  \n请求报文：%s  \n响应报文：%s   %s   \n| %sms  rid:%s", url, reqBody, resBody, errMsg, ms, rid)
	}()
	bodyBytes, _ := json.Marshal(bodyMap)
	reqBody = string(bodyBytes)
	resp, err := http.Post(url, "application/json", strings.NewReader(reqBody))
	if err != nil {
		err = errors.New("UP_HTTP_ERR:" + err.Error())
		return
	}
	if resp.StatusCode != http.StatusOK {
		err = errors.New(resp.Status)
		return
	}
	bytes, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return
	}
	resBody = string(bytes) // 银联有可能响应数据是错误网页，非JSON数据
	err = json.Unmarshal(bytes, &respBody)
	if err != nil || respBody == nil {
		err = errors.New("GCUP:响应数据异常 rid:" + rid)
		return
	}

	return
}

//Call RSASha256方式签名，其它签名或者加密方式请使用Post方法
func Call(c *Config, path string, bm *BodyMap, result interface{}) (err error) {
	//计算签名
	signature := bm.Sha256Sign(c.Secret)
	bm.Set("signature", signature)

	respBody, err := Post(c, path, bm)
	if err != nil {
		return
	}
	/*
		a10	不合法的backend_token，或已过期（参见6.1.1获取backendToken章节，重新获取backend_token）
		a20	不合法的frontend_token，或已过期（参见6.1.2获取frontToken章节，重新获取front_token）
		a31	不合法的授权code，或已过期（参见5.3系统对接步骤章节，参见常见问题解答）
		以上3个错误码需业务服务重新刷新
	*/
	switch respBody.Resp {
	case EA10.Code:
		return EA10
	case EA20.Code:
		return EA20
	case EA31.Code:
		return EA31
	}
	if respBody.Resp != E00.Code {
		return errors.New(respBody.Resp + ":" + respBody.Msg)
	}
	b, err := json.Marshal(respBody.Params)
	if err != nil {
		return
	}
	err = json.Unmarshal(b, result)
	if err != nil {
		return
	}
	return
}

type RespBody struct {
	Resp   string      `json:"resp"`
	Msg    string      `json:"msg"`
	Params interface{} `json:"params"`
}

//GetRandomString 获取随机字符串 length：字符串长度
func GetRandomString(length int) string {
	str := "0123456789AaBbCcDdEeFfGgHhIiJjKkLlMmNnOoPpQqRrSsTtUuVvWwXxYyZz"
	var (
		result []byte
		b      []byte
		r      *rand.Rand
	)
	b = []byte(str)
	r = rand.New(rand.NewSource(time.Now().UnixNano()))
	for i := 0; i < length; i++ {
		result = append(result, b[r.Intn(len(b))])
	}
	return string(result)
}

//Decode3DES 云闪付敏感数据解密
func Decode3DES(symmetricKey string, v string) (val string, err error) {
	if v == "" {
		val = ""
		return
	}
	bytes, err := hex.DecodeString(symmetricKey)
	if err != nil {
		return
	}
	key := string(bytes)
	val, err = TRIPLE_DES_ECB_PKCS5_Decode(v, key)
	return
}

//Encode3DES 云闪付敏感数据加密
func Encode3DES(symmetricKey string, v string) (val string, err error) {
	if v == "" {
		val = ""
		return
	}
	bytes, err := hex.DecodeString(symmetricKey)
	if err != nil {
		return
	}
	key := string(bytes)
	val, err = TRIPLE_DES_ECB_PKCS5_Encode(v, key)
	return
}
