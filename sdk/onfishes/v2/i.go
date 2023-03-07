package v2

import (
	"crypto/md5"
	"crypto/rand"
	"crypto/rsa"
	"crypto/x509"
	"encoding/base64"
	"encoding/json"
	"fmt"
	log "github.com/sirupsen/logrus"
	"io/ioutil"
	"math/big"
	"net/http"
	"strconv"
	"strings"
	"time"
)

/*
1.请求规则
接口默认请求方式为Post，请求数据格式为key=value键值方式，响应数据格式
统一为：json。
2.签名规则
对指定参数进行MD5运算操作，运算结果统一为大写32位字符串（不同接口参与
签名的参数可能不一样）。
3.前置资料
接口文档中：
URL（请求地址）
MemberAmountCode（扣款账户）
AppKey（账号标识）
AppSecret（账号密匙）
ProductCode（产品编号）
请从供货商相关业务负责人获取。
4.提交订单结果异常
订单提交接口请求响应非约定的合法数据以及响应超时等异常情况，不可做失败、
成功处理，请通过订单主动查询接口或客服核实确定订单状态。
5.订单状态同步
订单提交成功后20分钟后仍然没收到结果通知，请做主动查询处理。
订单提交15分钟后订单查询结果为“订单不存在”可做失败处理（查询订单号错
误情况除外）。
6.在文档中提到的“商户”、“会员”代指业务使用者。“服务商”、“供货商”、“服
务提供商”代指业务接口服务提供商。
7.接口中状态值或者充值状态（如：Code,OrderState）如果实际返数据异常或非文档描述
说明的状态请做异常处理并且向我方进行反馈。
*/

type Client struct {
	Cfg *Config
}

func NewClient(cfg *Config) *Client {
	return &Client{Cfg: cfg}
}

// 接口请求
func (c *Client) Request(path string, params string, result interface{}) (err error) {
	var (
		url          = c.Cfg.BaseServiceUrl + path
		plog         = log.WithField("channel", "onfishes").WithField("requestId", Rand32())
		requestBody  = params
		responseBody string
		begTime      = time.Now().UnixNano()
	)
	defer func() {
		endTime := time.Now().UnixNano()
		plog.Info("请求URL ", url)
		plog.WithField("requestBody", requestBody).Info("请求报文")
		plog.WithField("responseBody", responseBody).Info("响应报文,耗时", strconv.FormatInt((endTime-begTime)/1e6, 10), "ms")
		if err != nil {
			plog.Error("接口调用异常", err.Error())
		}
	}()
	resp, err := http.Post(url, "application/x-www-form-urlencoded", strings.NewReader(params))
	if err != nil {
		return
	}
	rbytes, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return
	}
	responseBody = string(rbytes)
	err = json.Unmarshal(rbytes, result)
	if err != nil {
		return
	}
	return
}

// 签名
func Md5Sign(v string) string {
	return strings.ToUpper(fmt.Sprintf("%x", md5.Sum([]byte(v))))
}

func Timestamp() string {
	return strconv.FormatInt(time.Now().Unix()*1000, 10)
}

// Rand32 使用crypto/rand 随机赋值byte数组， 然后md5返回32位十六进制字符串
func Rand32() string {
	var b = make([]byte, 48)
	if _, err := rand.Read(b); err != nil {
		return ""
	}
	return fmt.Sprintf("%x", md5.Sum(b))
}

// 私钥解密
func RsaDecrypt(val, priKey string) (plainText string, err error) {
	cipherText, err := base64.StdEncoding.DecodeString(val)
	if err != nil {
		return
	}
	priBytes, err := base64.StdEncoding.DecodeString(priKey)
	if err != nil {
		return
	}
	priv, err := x509.ParsePKCS8PrivateKey(priBytes)
	if err != nil {
		return
	}
	privateKey := priv.(*rsa.PrivateKey)

	// 解密 RSA/ECB/NoPadding
	c := new(big.Int).SetBytes(cipherText)
	plainText = string(c.Exp(c, privateKey.D, privateKey.N).Bytes())
	return
}
