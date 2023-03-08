package helper

import (
	"encoding/json"
	"fmt"
)

// QueryRefund 查询退款
// 接口文档地址 https://developer.open-douyin.com/docs/resource/zh-CN/mini-app/develop/server/trading/refund/query-refund
func QueryRefund(req *QueryRefundRequest) (res *QueryRefundResponse, err error) {
	url := QUERY_REFUND_URL
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

type QueryRefundRequest struct {
	RefundId    string `json:"refund_id,omitempty"`     //抖音开平内部交易退款单号
	OutRefundNo string `json:"out_refund_no,omitempty"` //开发者侧退款单号 `
}

type QueryRefundResponse struct {
	PubRes
	RespExtra struct {
		Logid string `json:"logid"`
	} `json:"resp_extra"`
	Data struct {
		RefundTotalAmount   int         `json:"refund_total_amount"` //退款金额，单位[分]
		RefundStatus        string      `json:"refund_status"`       //退款状态 退款中- PROCESSING 已退款- SUCCESS 退款失败- FAIL
		ItemOrderDetail     interface{} `json:"item_order_detail"`   //商品单信息 (交易系统订单退款才有的信息)
		MerchantAuditDetail struct {    //退款审核信息 (交易系统订单退款才有的信息)
			NeedRefundAudit int    `json:"need_refund_audit"`
			RefundAuditDead int64  `json:"refund_audit_dead"` //退款审核的最后期限，过期无需审核，自动退款，13 位 unix 时间戳，精度：毫秒
			AuditStatus     string `json:"audit_status"`      //退款审核状态： INIT：初始化 TOAUDIT：待审核 AGREE：同意 DENY：拒绝 OVERTIME：超时未审核自动同意
			DenyMessage     string `json:"deny_message"`      //不同意退款信息
		} `json:"merchant_audit_detail"`
		OrderID     string `json:"order_id"`      //系统订单信息，开放平台生成的订单号
		RefundID    string `json:"refund_id"`     //系统退款单号，开放平台生成的退款单号
		OutRefundNo string `json:"out_refund_no"` //开发者系统生成的退款单号，与抖音开平退款单号 refund_id 唯一关联
		RefundAt    int64  `json:"refund_at"`     //退款时间，13 位毫秒时间戳，只有已退款才有退款时间
		Message     string `json:"message"`       //退款结果信息，非商家拒绝退款导致的退款失败，可以通过该字段了解退款失败原因
	} `json:"data"`
}
