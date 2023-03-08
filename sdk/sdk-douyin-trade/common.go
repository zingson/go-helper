package helper

import (
	"bytes"
	"crypto"
	"crypto/rand"
	"crypto/rsa"
	"crypto/sha256"
	"crypto/x509"
	"encoding/base64"
	"encoding/json"
	"encoding/pem"
	"errors"
	"fmt"
	"github.com/sirupsen/logrus"
	douyinpoi "github.com/zingson/go-helper/sdk/sdk-douyin-poi"
	"io/ioutil"
	"math/big"
	"net/http"
	"strconv"
	"strings"
	"time"
)

// 通用错误码
type errorCode int64

const (
	// 通用状态码
	ERROR_CODE_0       errorCode = 0       // 成功
	ERROR_CODE_2100004 errorCode = 2100004 // 系统繁忙，此时请开发者稍候再试
	ERROR_CODE_2100005 errorCode = 2100005 // 参数不合法
	ERROR_CODE_2100007 errorCode = 2100007 // 无权限操作
	ERROR_CODE_2100009 errorCode = 2100009 // 用户被禁封使用该操作
	ERROR_CODE_2190001 errorCode = 2190001 // quota已用完
	ERROR_CODE_2190004 errorCode = 2190004 // 应用未获得该能力
	ERROR_CODE_2190015 errorCode = 2190015 // 请求参数access_token openid不匹配
	ERROR_CODE_2190016 errorCode = 2190016 // 当前应用已被封禁或下线
)

// 常用域名
const (
	SALT = "ERqy5LmZNGb8O8xqBQvYZefaWA3s7xUVmKfVtufQ"

	WEBSITE_URl           = "https://developer.toutiao.com"          // 域名
	GET_ACCESS_TOKEN_URL  = "/api/apps/v2/token"                     // getAccessToken
	ORDER_SYNC_URL        = "/api/apps/order/v2/push"                // 担保支付的订单同步
	ORDER_SETTLE_URL      = "/api/apps/ecpay/v1/settle"              // 担保支付的发起结算及分账
	MERCHANT_WITHDRAW_URL = "/api/apps/ecpay/saas/merchant_withdraw" // 商户提现

	SET_CALL_BACK_URL           = "/api/apps/trade/v2/settings"                // 设置回调地址
	QUERY_SETTINGS_URL          = "/api/apps/trade/v2/query_settings"          // 查询回调地址
	QUERY_ORDER_URL             = "/api/apps/trade/v2/query_order"             // 查询订单信息
	PUSH_DELIVERY_URL           = "/api/apps/trade/v2/push_delivery"           // 推送核销状态
	CREATE_SETTLE_URL           = "/api/apps/trade/v2/create_settle"           // 发起分账 一笔订单到达结算周期后，开发者可以通过分账接口将这笔订单产生的资金结算给各个分账方。
	CREATE_REFUND_URL           = "/api/apps/trade/v2/create_refund"           // 开发者发起退款
	MERCHANT_AUDIT_CALLBACK_URL = "/api/apps/trade/v2/merchant_audit_callback" // 同步退款审核结果
	QUERY_REFUND_URL            = "/api/apps/trade/v2/query_refund"            // 同步退款审核结果
	CREATE_ORDER_URL            = "/api/apps/trade/v2/create_order"            // 预下单
)

type DouyinConfig struct {
	AppID      string `json:"appId"`        // 应用appId
	AppKey     string `json:"clientKey"`    // 应用唯一标识
	AppSecret  string `json:"clientSecret"` // 应用唯一标识对应的密钥
	PrivateKey string `json:"privateKey"`   // 应用私钥
	PublicKey  string `json:"publicKey"`    // 平台公钥
}

// PubRes 响应公共参数
type PubRes struct {
	ErrNo   int64  `json:"err_no"`   //错误码，详见后文错误码列表，成功时为 0。
	ErrTips string `json:"err_tips"` //错误提示，这个字段会给出具体的错误原因。
}

// Request 发送接口请求
func Request(conf *douyinpoi.ClientTokenReq, url, method, reqBody string) (response string, err error) {
	var (
		rid    = RandomString(32)
		bmilli = time.Now().UnixMilli()
		reqUrl = WEBSITE_URl + url
	)
	defer func() {
		var errMsg string
		if err != nil {
			errMsg = ", 异常：" + err.Error() + " ,rid=" + rid
		}
		millisecond := fmt.Sprintf("%d", time.Now().UnixMilli()-bmilli)
		logrus.WithField("rid", rid).WithField("millisecond", millisecond).Infof("抖音交易系统请求：URL:%s, 请求报文：%s, 响应报文：%s %s  | %sms", reqUrl, reqBody, response, errMsg, millisecond)
	}()

	timestamp := strconv.FormatInt(time.Now().Unix(), 10)
	nonceStr := RandomString(32)
	signature, _ := GenSign(method, url, timestamp, nonceStr, reqBody, conf.PrivateKey)
	authorization := GenByteAuthorization(conf.AppId, timestamp, nonceStr, signature)
	request, err := http.NewRequest(method, reqUrl, strings.NewReader(reqBody))
	if err != nil {
		return
	}
	request.Header.Set("Content-Type", "application/json")
	request.Header.Set("Accept", "application/json")
	request.Header.Set("Byte-Authorization", authorization)
	defer request.Body.Close()
	resp, err := http.DefaultClient.Do(request)
	if err != nil {
		return
	}
	defer resp.Body.Close()
	resBytes, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return
	}
	response = string(resBytes)
	// HTTP 返回非200，直接返回错误
	if resp.StatusCode != 200 {
		var eres *PubRes
		json.Unmarshal(resBytes, &eres)
		err = errors.New(strconv.FormatInt(eres.ErrNo, 10) + ":" + eres.ErrTips)
		return
	}

	//抖音返回数据验签
	resTimestamp := resp.Header.Get("Byte-Timestamp")
	resNonce := resp.Header.Get("Byte-Nonce-Str")
	resSignature := resp.Header.Get("Byte-Signature")
	if resSignature == "" || resNonce == "" || resTimestamp == "" {
		return
	}
	ok, err := CheckSign(resTimestamp, resNonce, response, resSignature, conf.PublicKey)
	if err != nil {
		return
	}
	if !ok {
		return response, errors.New("签名校验失败")
	}
	return
}

func GenByteAuthorization(appid, timestamp, nonceStr, signature string) string {
	authType := "SHA256-RSA2048" //固定字符串
	return fmt.Sprintf(`%s appid="%s",nonce_str="%s",timestamp="%s",key_version="2",signature="%s"`, authType, appid, nonceStr, timestamp, signature)
}

// GenSign 生成签名
func GenSign(method, url, timestamp, nonce, body, priKeyStr string) (string, error) {
	priKey, err := PemToRSAPriKey(priKeyStr)
	if err != nil {
		return "", err
	}
	//组装被加密的字符串
	targetStr := method + "\n" + url + "\n" + timestamp + "\n" + nonce + "\n" + body + "\n"
	//加密
	h := sha256.New()
	h.Write([]byte(targetStr))
	digestBytes := h.Sum(nil)
	signBytes, err := rsa.SignPKCS1v15(rand.Reader, priKey, crypto.SHA256, digestBytes)
	if err != nil {
		return "", err
	}
	sign := base64.StdEncoding.EncodeToString(signBytes)
	return sign, nil
}

func PemToRSAPriKey(pemKeyStr string) (pri *rsa.PrivateKey, err error) {
	keyBytes := []byte(pemKeyStr)
	block, _ := pem.Decode(keyBytes)
	if block == nil {
		err = errors.New("rsa private key error")
		return
	}
	privateKey, err := x509.ParsePKCS8PrivateKey(block.Bytes)
	if err != nil {
		return
	}
	pri = privateKey.(*rsa.PrivateKey)
	return
}

// CheckSign  签名验证
func CheckSign(timestamp, nonce, body, signature, pubKeyStr string) (bool, error) {

	pubKey, err := PemToRSAPublicKey(pubKeyStr)
	if err != nil {
		return false, err
	}

	hashed := sha256.Sum256([]byte(timestamp + "\n" + nonce + "\n" + body + "\n"))
	signBytes, err := base64.StdEncoding.DecodeString(signature)
	if err != nil {
		return false, err
	}
	err = rsa.VerifyPKCS1v15(pubKey, crypto.SHA256, hashed[:], signBytes)
	return err == nil, nil
}

func PemToRSAPublicKey(pemKeyStr string) (*rsa.PublicKey, error) {
	block, _ := pem.Decode([]byte(pemKeyStr))
	if block == nil || len(block.Bytes) == 0 {
		return nil, fmt.Errorf("empty block in pem string")
	}
	key, err := x509.ParsePKIXPublicKey(block.Bytes)
	if err != nil {
		return nil, err
	}
	switch key := key.(type) {
	case *rsa.PublicKey:
		return key, nil
	default:
		return nil, fmt.Errorf("not rsa public key")
	}
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
