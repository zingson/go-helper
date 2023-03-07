package v3

// Query 查询订单API
// 文档：https://pay.weixin.qq.com/wiki/doc/apiv3/wxpay/pay/transactions/chapter3_5.shtml
func (pay *PayService) Query(transactionId, outTradeNo string) (r *PayResult, err error) {
	// 微信支付订单号查询
	if transactionId != "" {
		err = pay.HttpGet("/v3/pay/transactions/id/"+transactionId+"?mchid="+pay.Mchid, &r)
		return
	}
	//商户订单号查询
	err = pay.HttpGet("/v3/pay/transactions/out-trade-no/"+outTradeNo+"?mchid="+pay.Mchid, &r)
	return
}

type PayResult struct {
	Appid          string     `json:"appid"`            // APPID
	Mchid          string     `json:"mchid"`            // 商户号
	OutTradeNo     string     `json:"out_trade_no"`     // 商户平台订单号
	TransactionId  string     `json:"transaction_id"`   // 微信支付交易号
	TradeType      TradeType  `json:"trade_type"`       // 交易类型
	TradeState     TradeState `json:"trade_state"`      // 交易状态
	TradeStateDesc string     `json:"trade_state_desc"` // 交易状态描述
	BankType       string     `json:"bank_type"`        // 银行类型  银行对照表：https://pay.weixin.qq.com/wiki/doc/apiv3/terms_definition/chapter1_1_3.shtml#part-6
	Attach         string     `json:"attach"`           // 附加数据，原样返回
	SuccessTime    Time       `json:"success_time"`     //支付完成时间，遵循rfc3339标准格式
	Payer          PayPayer   `json:"payer"`            //支付者信息
	Amount         PayAmount  `json:"amount"`           //订单金额信息，当支付成功时返回该字段。
	//PromotionDetail interface{} `json:"promotion_detail"`
}
