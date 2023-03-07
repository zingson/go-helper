package banknbcb

import "encoding/json"

// QueryOrderStatus 查询订单
func QueryOrderStatus(c *Config, trace_no, out_trade_id string) (r *QueryOrderStatusResult, err error) {
	p := make(map[string]interface{})
	p["tran_code"] = "queryOrderStatus"
	p["mchnt_cd"] = c.MchntCd
	p["trace_no"] = trace_no         //商户或终端自 行生成的订单 流水号
	p["out_trade_id"] = out_trade_id //易收宝生成的 订单流水号
	err = POST(c, p, &r)
	return
}

type QueryOrderStatusResult struct {
	TotalFee        string      `json:"total_fee"`         // 交易金额，单位分
	RefundAmount    string      `json:"refundAmount"`      // 已退金额
	TransStatus     TransStatus `json:"trans_status"`      //交易状态
	TransStatusDesc string      `json:"trans_status_desc"` //
	FundType        string      `json:"fund_type"`         // 0=余额  1=借记卡  2=贷记卡
	TraceNo         string      `json:"trace_no"`          //商户或终端自行生成 的订单流水号
	OutTradeId      string      `json:"out_trade_id"`      //易收宝生成的订单流 水号
	TransDatetime   string      `json:"trans_datetime"`    //交易时间，14位
	ChannelType     string      `json:"channel_type"`      // 1=微信 2=支付宝 3=银联二维码
	Remark          string      `json:"remark"`            //备注
	TradeId         string      `json:"trade_id"`          //银联交易流水
	HanglingCharge  string      `json:"hanglingCharge"`    //订单手续费
}

func (o *QueryOrderStatusResult) String() (s string) {
	b, _ := json.Marshal(o)
	return string(b)
}
