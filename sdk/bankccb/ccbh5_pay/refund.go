package ccbh5_pay

import (
	"encoding/json"
	"errors"
	"fmt"
	"time"
)

//Refund svc_occPlatRefund 订单退款
func Refund(conf *Config, orderId string, refundAmt int64, payTime string) (body *RefundBody, err error) {
	layout := "20060102150405"
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

	cldBody, err := Post(conf, conf.ServiceOccplatreq, "svc_occPlatRefund", map[string]string{
		"CUSTOMERID": conf.MerchantId,
		"BRANCHID":   conf.BranchId,
		"MONEY":      fmt.Sprintf("%.2f", float64(refundAmt)/100), //退款金额 单位：元
		"ORDER":      orderId,                                     //服务方订单号
		"STDT_TM":    stdtTm,                                      //开始日期时间
		"EDDT_TM":    eddtTm,                                      //结束日期时间
	})
	if err != nil {
		return
	}
	var resResult RefundResponse
	err = json.Unmarshal([]byte(cldBody), &resResult)
	if err != nil {
		return
	}
	if resResult.CLD_HEADER.CLD_TX_RESP.CLD_CODE != "CLD_SUCCESS" { // 判断成功状态码
		err = errors.New(resResult.CLD_HEADER.CLD_TX_RESP.CLD_DESC)
		return
	}
	body = resResult.CLD_BODY
	return
}

type RefundBody struct {
	ORDER_NUM  string `json:"ORDER_NUM"`  //订单号
	PAY_AMOUNT string `json:"PAY_AMOUNT"` //支付金额 元
	AMOUNT     string `json:"AMOUNT"`     //退款金额 元
}

/*
CLD_HEADER
...CLD_TX_CHNL
...CLD_TX_TIME
...CLD_TX_CODE
...CLD_TX_SEQ
...CLD_TX_RESP
......CLD_CODE
......CLD_DESC
*/
type RefundResponse struct {
	CLD_HEADER struct {
		CLD_TX_CHNL string `json:"CLD_TX_CHNL"`
		CLD_TX_TIME string `json:"CLD_TX_TIME"`
		CLD_TX_CODE string `json:"CLD_TX_CODE"`
		CLD_TX_SEQ  string `json:"CLD_TX_SEQ"`
		CLD_TX_RESP struct {
			CLD_CODE string `json:"CLD_CODE"` // 响应码
			CLD_DESC string `json:"CLD_DESC"`
		} `json:"CLD_TX_RESP"`
	} `json:"CLD_HEADER"`
	CLD_BODY *RefundBody `json:"CLD_BODY"`
}
