package v3

//Refund 申请退款API
func (pay *PayService) Refund(params *RefundParams) (r *RefundResult, err error) {
	err = pay.HttpPost("/v3/refund/domestic/refunds", params, &r)
	return
}

type RefundParams struct {
	TransactionId string               `json:"transaction_id"`
	OutTradeNo    string               `json:"out_trade_no"`
	OutRefundNo   string               `json:"out_refund_no"`
	Reason        string               `json:"reason,omitempty"`
	NotifyUrl     string               `json:"notify_url,omitempty"`
	FundsAccount  string               `json:"funds_account,omitempty"`
	Amount        *RefundParamsAmount  `json:"amount"`
	GoodsDetail   []*RefundParamsGoods `json:"goods_detail,omitempty"`
}

type RefundParamsAmount struct {
	Refund   int64  `json:"refund"`
	Total    int64  `json:"total"`
	Currency string `json:"currency,omitempty"`
}

type RefundParamsGoods struct {
	MerchantGoodsId  string `json:"merchant_goods_id"`
	WechatpayGoodsId string `json:"wechatpay_goods_id"`
	GoodsName        string `json:"goods_name"`
	UnitPrice        int64  `json:"unit_price"`
	RefundAmount     int64  `json:"refund_amount"`
	RefundQuantity   int64  `json:"refund_quantity"`
}

type RefundResult struct {
	RefundId            string                `json:"refund_id"`
	OutRefundNo         string                `json:"out_refund_no"`
	TransactionId       string                `json:"transaction_id"`
	OutTradeNo          string                `json:"out_trade_no"`
	Channel             string                `json:"channel"`
	UserReceivedAccount string                `json:"user_received_account"`
	SuccessTime         string                `json:"success_time"`
	CreateTime          string                `json:"create_time"`
	Status              string                `json:"status"`
	FundsAccount        string                `json:"funds_account"`
	Amount              *RefundResultAmount   `json:"amount"`
	PromotionDetail     []*RefundResultDetail `json:"promotion_detail"`
}

type RefundResultAmount struct {
	Total            int64  `json:"total"`
	Refund           int64  `json:"refund"`
	PayerTotal       int64  `json:"payer_total"`
	PayerRefund      int64  `json:"payer_refund"`
	SettlementRefund int64  `json:"settlement_refund"`
	SettlementTotal  int64  `json:"settlement_total"`
	DiscountRefund   int64  `json:"discount_refund"`
	Currency         string `json:"currency"`
}

type RefundResultDetail struct {
	PromotionId  string                     `json:"promotion_id"`
	Scope        string                     `json:"scope"`
	Type         string                     `json:"type"`
	Amount       int64                      `json:"amount"`
	RefundAmount int64                      `json:"refund_amount"`
	GoodsDetail  []*RefundResultDetailGoods `json:"goods_detail,omitempty"`
}

type RefundResultDetailGoods struct {
	MerchantGoodsId  string `json:"merchant_goods_id"`
	WechatpayGoodsId string `json:"wechatpay_goods_id"`
	GoodsName        string `json:"goods_name"`
	UnitPrice        int64  `json:"unit_price"`
	RefundAmount     int64  `json:"refund_amount"`
	RefundQuantity   int64  `json:"refund_quantity"`
}
