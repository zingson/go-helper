package helper

import (
	"encoding/json"
	"fmt"
)

// CreateRefund 发起退款
// 接口文档地址 https://developer.open-douyin.com/docs/resource/zh-CN/mini-app/develop/server/trading/refund/apply
func CreateRefund(req *CreateRefundRequest) (res *CreateRefundResponse, err error) {
	url := CREATE_REFUND_URL
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

type CreateRefundRequest struct {
	OutOrderNo        string             `json:"out_order_no"`         //开发者侧订单 id
	OutRefundNo       string             `json:"out_refund_no"`        //开发者侧退款单号
	CpExtra           string             `json:"cp_extra,omitempty"`   //开发者自定义透传字段
	OrderEntrySchema  *OrderEntrySchema  `json:"order_entry_schema"`   //退款单的跳转的 schema
	NotifyUrl         string             `json:"notify_url,omitempty"` //退款结果通知地址，必须是 HTTPS 类型，若不填，默认使用在https://developer.open-douyin.com/microapp/ttf4d2826f6becc24001/pay页面设置的支付回调地址。 ttf4d2826f6becc24001为小程序的 AppID。
	ItemOrderDetail   []*ItemOrderDetail `json:"item_order_detail,omitempty"`
	RefundTotalAmount int64              `json:"refund_total_amount,omitempty"`
}

type OrderEntrySchema struct {
	Path   string `json:"path"`             //订单详情页路径，没有前导的/，该字段不能为空，长度 <= 512byte pages/xxxindexxxx
	Params string `json:"params,omitempty"` //路径参数，自定义的 json 结构，序列化成字符串存入该字段，平台不限制，但是写入的内容需要能够保证生成访问订单详情的 schema 能正确跳转到小程序内部的订单详情页，长度 <= 512byte  {\"id\":\"xxxxxx\"}
}

type ItemOrderDetail struct {
	ItemOrderId  string `json:"item_order_id"`           //商品单号
	RefundAmount int64  `json:"refund_amount,omitempty"` //该 item_order 需要退款的金额，单位分，不能大于该 item_order 实付金额且要大于 0
}

type CreateRefundResponse struct {
	PubRes
	RespExtra struct {
		Logid string `json:"logid"`
	} `json:"resp_extra"`
	Data struct {
		RefundId            string `json:"refund_id"`
		RefundAuditDeadline string `json:"refund_audit_deadline"`
	} `json:"data"`
}
