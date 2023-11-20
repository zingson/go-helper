package wxapp

// UploadShippingInfo 发货信息录入接口
// https://developers.weixin.qq.com/miniprogram/dev/platform-capabilities/business-capabilities/order-shipping/order-shipping.html
func UploadShippingInfo(accessToken string, params *UploadShippingInfoParams) (err error) {
	_, err = post[Result]("https://api.weixin.qq.com/wxa/sec/order/upload_shipping_info?access_token="+accessToken, params)
	return
}

type UploadShippingInfoParams struct {
	OrderKey      OrderKey        `json:"order_key"`
	LogisticsType int64           `json:"logistics_type"`
	DeliveryMode  int64           `json:"delivery_mode"`
	ShippingList  []*ShippingItem `json:"shipping_list"`
	UploadTime    string          `json:"upload_time"`
	Payer         Payer           `json:"payer"`
}

type OrderKey struct {
	OrderNumberType int64  `json:"order_number_type"`
	TransactionId   string `json:"transaction_id"`
	Mchid           string `json:"mchid"`
	OutTradeNo      string `json:"out_trade_no"`
}

type ShippingItem struct {
	ItemDesc string `json:"item_desc"`
}

type Payer struct {
	Openid string `json:"openid"`
}
