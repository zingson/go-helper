package banknbcb

import "encoding/json"

func Refund(c *Config, params *RefundParams) (r *RefundResult, err error) {
	if params.StaffId == "" {
		params.StaffId = "0000"
	}
	b0, err := json.Marshal(params)
	if err != nil {
		return
	}
	p := make(map[string]interface{})
	err = json.Unmarshal(b0, &p)
	if err != nil {
		return
	}
	p["tran_code"] = "refund"
	err = POST(c, p, &r)
	return
}

type RefundParams struct {
	StaffId        string `json:"staff_id"`          //商户下登记的 员工号。
	RefundFee      string `json:"refund_fee"`        //退款金额 单位分
	TraceNo        string `json:"trace_no"`          //商户或终端自行生成的订单 流水号。
	OrigTraceNo    string `json:"orig_trace_no"`     //商户原交易订 单流水号。
	OrigOutTradeId string `json:"orig_out_trade_id"` //原交易订单易 收宝系统订单 号
	Remark         string `json:"remark"`            // 备注
}

type RefundResult struct {
	TotalFee       string `json:"total_fee"`
	FundType       string `json:"fund_type"`
	RefundFee      string `json:"refund_fee"`
	OutTradeId     string `json:"out_trade_id"`
	TransDatetime  string `json:"trans_datetime"`  //t14
	RefundDatetime string `json:"refund_datetime"` //t14
	HanglingCharge string `json:"hangling_charge"`
}
