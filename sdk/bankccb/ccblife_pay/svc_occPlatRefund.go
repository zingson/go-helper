package ccblife_pay

import (
	"fmt"
	"time"
)

const layout = "20060102150405"

// PlatRefund 订单退款
func PlatRefund(conf *Config, orderId string, refundAmt int64, payTime string) (body RefundBody, err error) {
	pt, err := time.Parse(layout, payTime)
	if err != nil {
		return
	}
	stdtTm := pt.Add(-4 * time.Hour).Format(layout)
	eddtTm := pt.Add(4 * time.Hour).Format(layout)
	nowTm := time.Now().Local().Format(layout)
	if eddtTm > nowTm {
		eddtTm = nowTm[:8] + eddtTm[8:]
	}

	bldBody := map[string]string{
		"CUSTOMERID": conf.MerchantId,                             //
		"BRANCHID":   conf.BranchId,                               //
		"MONEY":      fmt.Sprintf("%.2f", float64(refundAmt)/100), //退款金额 单位：元
		"ORDER":      orderId,                                     //服务方订单号
		"STDT_TM":    stdtTm,                                      //开始日期时间
		"EDDT_TM":    eddtTm,                                      //结束日期时间
	}
	return Call[RefundBody](conf, conf.ServiceOccplatreq, "svc_occPlatRefund", bldBody)
}

type RefundBody struct {
	OrderNum  string `json:"ORDER_NUM"`  //订单号
	PayAmount string `json:"PAY_AMOUNT"` //支付金额 元
	Amount    string `json:"AMOUNT"`     //退款金额 元
}
