package v3

import (
	"encoding/json"
	"errors"
	"fmt"
	"github.com/sirupsen/logrus"
	"io"
	"net/http"
	"time"
)

// Notify 微信支付通知
func (pay *PayService) Notify(request *http.Request, writer http.ResponseWriter, resolve func(rs *PayResult) error) {
	var (
		err       error
		rid       = RandomString(32)
		milli     = time.Now().UnixMilli()
		notifyURI = request.RequestURI
		oriBody   string
		reqBody   string
		resBody   string
		mchid     string
	)
	defer func() {
		resBody = notifyResponse(err)
		logrus.
			WithField("rid", rid).
			WithField("millisecond", fmt.Sprintf("%d", time.Now().UnixMilli()-milli)).
			Infof("wxpay mchid=%s 微信支付通知 \n通知URI：%s  \n通知明文：%s   \n通知密文：%s   \n响应报文：%s ", mchid, notifyURI, reqBody, oriBody, resBody)
		_, _ = writer.Write([]byte(resBody))
		if err != nil {
			writer.WriteHeader(500)
		}
	}()
	rBytes, err := io.ReadAll(request.Body)
	if err != nil {
		return
	}
	oriBody = string(rBytes)

	// 验签
	err = NotifyVerifySign(pay.Client, request.Header, oriBody)
	if err != nil {
		return
	}

	// 解密
	reqBody, err = NotifyDecrypt(pay.Client, rBytes)
	if err != nil {
		return
	}

	//解析数据
	var payResult *PayResult
	err = json.Unmarshal([]byte(reqBody), &payResult)
	if err != nil {
		return
	}

	// 回调业务函数
	err = resolve(payResult)
	if err != nil {
		return
	}
	mchid = payResult.Mchid
}

// 通知应答
func notifyResponse(err error) string {
	if err != nil {
		return fmt.Sprintf(`{"code": "FAIL","message": "%s"}`, err.Error())
	}
	return `{"code": "SUCCESS","message": "成功"}`
}

// NotifyVerifySign 微信支付与退款通知验签
func NotifyVerifySign(c *Client, h http.Header, body string) (err error) {
	signature := h.Get("Wechatpay-Signature")
	serial := h.Get("Wechatpay-Serial")
	timestamp := h.Get("Wechatpay-Timestamp")
	nonce := h.Get("Wechatpay-Nonce")

	pubKey, err := c.GetWxPubKey(serial)
	if err != nil {
		return
	}
	ok, err := V3SignVery(signature, timestamp, nonce, body, pubKey)
	if err != nil {
		return
	}
	if !ok {
		err = errors.New("WX_NOTIFY:签名校验失败")
		return
	}
	return
}

type NotifyContent struct {
	Id           string                 `json:"id"`            //通知的唯一ID
	CreateTime   string                 `json:"create_time"`   //通知创建的时间
	EventType    string                 `json:"event_type"`    //通知的类型，支付成功通知的类型为TRANSACTION.SUCCESS
	ResourceType string                 `json:"resource_type"` //通知的资源数据类型，支付成功通知为encrypt-resource
	Summary      string                 `json:"summary"`       //回调摘要 示例值：支付成功
	Resource     *NotifyContentResource `json:"resource"`
}

type NotifyContentResource struct {
	Algorithm      string `json:"algorithm"`       //对开启结果数据进行加密的加密算法，目前只支持AEAD_AES_256_GCM
	Ciphertext     string `json:"ciphertext"`      //Base64编码后的开启/停用结果数据密文
	AssociatedData string `json:"associated_data"` //附加数据
	OriginalType   string `json:"original_type"`   //原始回调类型，为transaction
	Nonce          string `json:"nonce"`           //加密使用的随机串
}

// NotifyDecrypt 解析通知内容
func NotifyDecrypt(c *Client, body []byte) (plaintext string, err error) {
	var nc *NotifyContent
	err = json.Unmarshal(body, &nc)
	if err != nil {
		return
	}
	// 解密
	return AesGcmDecrypt(nc.Resource.Ciphertext, nc.Resource.Nonce, nc.Resource.AssociatedData, c.V3Secret)
}
