package helper

import (
	"encoding/json"
	"fmt"
)

// MerchantAuditCallback 同步退款审核结果
// 接口文档地址 https://developer.open-douyin.com/docs/resource/zh-CN/mini-app/develop/server/trading/refund/audit-callback
func MerchantAuditCallback(req *MerchantAuditCallbackRequest) (res *PubRes, err error) {
	url := MERCHANT_AUDIT_CALLBACK_URL
	reqJson, err := json.Marshal(req)
	if err != nil {
		fmt.Printf("json marshal err:%v", err)
	}
	result, err := Request(nil, url, "POST", string(reqJson))
	if err = json.Unmarshal([]byte(result), &res); err != nil {
		fmt.Printf("json unmarshal err:%v：", err)
	}
	return
}

type MerchantAuditCallbackRequest struct {
	OutRefundNo       string `json:"out_refund_no"`          //开发者侧订单 id
	RefundAuditStatus int8   `json:"refund_audit_status"`    //审核状态 1：同意退款 2：不同意退款
	DenyMessage       string `json:"deny_message,omitempty"` //不同意退款信息(不同意退款时必填)
}
