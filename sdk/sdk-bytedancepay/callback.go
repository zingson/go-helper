package bytedancepay

import (
	"encoding/json"
	"errors"
	"github.com/sirupsen/logrus"
)

//Callback 回调通知 POST Content-Type: application/json'
// token 从小程序支付获取
// uri  当前请求uri
// reqBody 请求报文,json格式字符串
// out  输出
func Callback(token, uri string, reqBody string, resolve func(notify *Notify) error) (resBody string) {
	var err error
	var notify *Notify
	defer func() {
		errMsg := ""
		if err != nil {
			errMsg = "ZJCBERR:" + err.Error()
			resBody = "{\"err_no\":500,\"err_tips\":\"" + err.Error() + "\"}"
		}
		logrus.WithField("url", uri).Infof("ZJ 支付回调通知 POST %s  请求报文：%s  响应报文：%s  %s ", uri, reqBody, resBody, errMsg)
	}()
	err = json.Unmarshal([]byte(reqBody), &notify)
	if err != nil {
		return
	}
	//sign := CbSign(structToMap(notify), token)
	sign := CbSignNew(token, notify.Timestamp, notify.Nonce, notify.Msg)
	if sign != notify.MsgSignature {
		err = errors.New("ZJ_SIGN_ERR:支付回调签名验证错误")
		return
	}
	err = resolve(notify)
	if err != nil {
		return
	}
	resBody = "{\"err_no\":0,\"err_tips\":\"success\"}"
	return
}

type Notify struct {
	Timestamp    string     `json:"timestamp"`     // Unix 时间戳，10 位，整型数
	Nonce        string     `json:"nonce"`         //
	Msg          string     `json:"msg"`           //订单信息的 json 字符串
	MsgSignature string     `json:"msg_signature"` //签名
	Type         NotifyType `json:"type"`          // payment=支付成功通知  refund=退款通知
}

type NotifyType string

const (
	NOTIFY_TYPE_PAYMENT NotifyType = "payment"
	NOTIFY_TYPE_REFUND  NotifyType = "refund"
)

func (m *Notify) Payment() (p *PaymentNotify, err error) {
	if m.Type != "payment" {
		err = errors.New("非支付成功通知")
		return
	}
	err = json.Unmarshal([]byte(m.Msg), &p)
	if err != nil {
		return
	}
	return
}

//Refund 退款通知参数
func (m *Notify) Refund() (p *RefundNotify, err error) {
	if m.Type != "refund" {
		err = errors.New("非退款通知")
		return
	}
	err = json.Unmarshal([]byte(m.Msg), &p)
	if err != nil {
		return
	}
	return
}

type PaymentNotify struct {
	Appid          string      `json:"appid"`
	CpOrderno      string      `json:"cp_orderno"`       //开发者传入订单号
	CpExtra        string      `json:"cp_extra"`         //预下单时开发者传入字段
	Way            string      `json:"way"`              //way 字段中标识了支付渠道：2-支付宝，1-微信
	PaymentOrderNo string      `json:"payment_order_no"` //
	TotalAmount    int         `json:"total_amount"`     //
	Status         OrderStatus `json:"status"`           //PROCESSING-处理中|SUCCESS-成功|FAIL-失败|TIMEOUT-超时
	SellerUid      string      `json:"seller_uid"`       //该笔交易卖家商户号
	Extra          string      `json:"extra"`            //该笔交易附加业务逻辑说明，例如 CPS 交易
	ItemId         string      `json:"item_id"`          //订单来源视频对应视频 id
}

type RefundNotify struct {
	Appid        string `json:"appid"`
	CpRefundno   string `json:"cp_refundno"`
	CpExtra      string `json:"cp_extra"`
	Status       string `json:"status"`
	RefundAmount int    `json:"refund_amount"`
}
