package banknbcb

import "encoding/json"

//QueryRefund 退款查询
func QueryRefund(c *Config, params *QueryRefundParams) (r *QueryRefundResult, err error) {
	b0, err := json.Marshal(params)
	if err != nil {
		return
	}
	p := make(map[string]interface{})
	err = json.Unmarshal(b0, &p)
	if err != nil {
		return
	}
	p["tran_code"] = "queryRefund"
	err = POST(c, p, &r)
	return
}

type QueryRefundParams struct {
	TraceNo    string `json:"trace_no"`
	OutTradeId string `json:"out_trade_id"`
}

type QueryRefundResult struct {
	RefundNo          string       `json:"refund_no"`
	TotalFee          string       `json:"total_fee"`
	RefundFee         string       `json:"refund_fee"`
	RefundAmount      string       `json:"refundAmount"`
	HanglingCharge    string       `json:"hanglingCharge"`
	OrigOutTradeId    string       `json:"orig_out_trade_id"`
	RefundStatus      RefundStatus `json:"refund_status"` //02 – 退款成功 11 – 退款中 12 – 退款错误
	RefundStatusDesc  string       `json:"refund_status_desc"`
	TradeType         string       `json:"trade_type"`
	RefundDatetime    string       `json:"refund_datetime"`
	OrigTransDatetime string       `json:"orig_trans_datetime"`
	Remark            string       `json:"remark"`
}
