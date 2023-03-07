package v3

import (
	"encoding/json"
	"strconv"
	"time"
)

// Jsapi 商户直连： JSAPI/小程序下单API 文档：https://pay.weixin.qq.com/wiki/doc/apiv3/apis/chapter3_1_1.shtml
func (pay *PayService) Jsapi(params *JsapiPay) (p *JsapiPayParams, err error) {

	// 微信下单
	var r *JsapiPayResult
	err = pay.HttpPost("/v3/pay/transactions/jsapi", params, &r)
	if err != nil {
		return
	}

	// 前端调起支付参数签名
	p = &JsapiPayParams{
		AppId:     params.Appid,
		TimeStamp: strconv.FormatInt(time.Now().Unix(), 10),
		NonceStr:  RandomString(32),
		Package:   "prepay_id=" + r.PrepayId,
		SignType:  "RSA",
		PaySign:   "",
	}
	s := p.AppId + "\n" + p.TimeStamp + "\n" + p.NonceStr + "\n" + p.Package + "\n"
	sign, err := RsaSignWithSha256(s, pay.mchPriKey)
	if err != nil {
		return
	}
	p.PaySign = sign

	return
}

type JsapiPay struct {
	Appid       string      `json:"appid"`                 //直连商户申请的公众号或者小程序appid。
	Mchid       string      `json:"mchid"`                 //直连商户的商户号，由微信支付生成并下发。
	Description string      `json:"description"`           //商品描述
	OutTradeNo  string      `json:"out_trade_no"`          //商户系统内部订单号
	TimeExpire  string      `json:"time_expire,omitempty"` //非必填。订单失效时间，遵循rfc3339标准格式，格式为YYYY-MM-DDTHH:mm:ss+TIMEZONE，
	Attach      string      `json:"attach,omitempty"`      //非必填。附加数据，在查询API和支付通知中原样返回，可作为自定义参数使用，有长度限制
	NotifyUrl   string      `json:"notify_url"`            //通知URL必须为直接可访问的URL，不允许携带查询串，要求必须为https地址。
	GoodsTag    string      `json:"goods_tag,omitempty"`   //非必填。订单优惠标记，有长度限制
	Amount      *PayAmount  `json:"amount"`                //订单金额信息
	Payer       *PayPayer   `json:"payer"`                 //支付者信息
	Detail      *PayDetail  `json:"detail,omitempty"`      //非必填。优惠功能,正常支付订单不用此参数
	SceneInfo   interface{} `json:"scene_info,omitempty"`  //非必填。支付场景描述，查看文档文档
	SettleInfo  interface{} `json:"settle_info,omitempty"` //非必填。结算信息，查看文档文档
}

type PayAmount struct {
	Total    int64  `json:"total"`              //订单总金额，单位为分。
	Currency string `json:"currency,omitempty"` //非必填。 CNY：人民币，境内商户号仅支持人民币。
}

type PayPayer struct {
	Openid string `json:"openid"` //用户在直连商户appid下的唯一标识。
}

type PayDetail struct {
	CostPrice   int64                  `json:"cost_price"`             //订单原价
	InvoiceId   string                 `json:"invoice_id,omitempty"`   //非必填。商品小票ID
	GoodsDetail []*JsapiPayDetailGoods `json:"goods_detail,omitempty"` //非必填。单品列表信息，条目个数限制：【1，6000】
}

type JsapiPayDetailGoods struct {
	MerchantGoodsId  string `json:"merchant_goods_id"`            //商户侧商品编码
	WechatpayGoodsId string `json:"wechatpay_goods_id,omitempty"` //非必填。微信侧商品编码
	GoodsName        string `json:"goods_name,omitempty"`         //非必填。商品名称
	Quantity         int64  `json:"quantity"`                     //商品数量
	UnitPrice        int64  `json:"unit_price"`                   //商品单价
}

type JsapiPayResult struct {
	PrepayId string `json:"prepay_id"` //预支付交易会话标识。用于后续接口调用中使用，该值有效期为2小时
}

type JsapiPayParams struct {
	AppId     string `json:"appId"`
	TimeStamp string `json:"timeStamp"`
	NonceStr  string `json:"nonceStr"`
	Package   string `json:"package"`
	SignType  string `json:"signType"`
	PaySign   string `json:"paySign"`
}

func (p *JsapiPayParams) JSON() string {
	v, _ := json.Marshal(p)
	return string(v)
}
