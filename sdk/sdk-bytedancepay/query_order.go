package bytedancepay

import (
	"encoding/json"
	"errors"
	"fmt"
)

//QueryOrder 订单查询  文档：https://microapp.bytedance.com/docs/zh-CN/mini-app/develop/server/ecpay/server-doc
func QueryOrder(conf *Config, params *QueryOrderParams) (rs *QueryOrderResult, err error) {
	pmap := structToMap(params)
	pmap["sign"] = getSign(pmap, conf.Salt)
	rBytes, err := json.Marshal(pmap)
	if err != nil {
		return
	}
	resBody, err := post(conf, "/api/apps/ecpay/v1/query_order", string(rBytes))
	if err != nil {
		return
	}
	err = json.Unmarshal([]byte(resBody), &rs)
	if err != nil {
		return
	}
	if rs.ErrNo != 0 {
		err = errors.New(fmt.Sprintf("Z%d:%s", rs.ErrNo, rs.ErrTips))
		return
	}
	return
}

type QueryOrderParams struct {
	AppId      string `json:"app_id"`
	OutOrderNo string `json:"out_order_no"` // 开发者侧的订单号, 同一小程序下不可重复
}

type QueryOrderResult struct {
	ErrNo       int                          `json:"err_no"`
	ErrTips     string                       `json:"err_tips"`
	OutOrderNo  string                       `json:"out_order_no"`
	OrderId     string                       `json:"order_id"`
	PaymentInfo *QueryOrderResultPaymentInfo `json:"payment_info"`
}

func (m *QueryOrderResult) Json() string {
	v, _ := json.Marshal(m)
	return string(v)
}

type QueryOrderResultPaymentInfo struct {
	TotalFee         int    `json:"total_fee"`
	OrderStatus      string `json:"order_status"`       //PROCESSING-处理中|SUCCESS-成功|FAIL-失败|TIMEOUT-超时
	PayTime          string `json:"pay_time"`           //支付时间
	Way              int    `json:"way"`                //way 字段中标识了支付渠道：2-支付宝，1-微信
	ChannelNo        string `json:"channel_no"`         //
	ChannelGatewayNo string `json:"channel_gateway_no"` //
	SellerUid        string `json:"seller_uid"`         //该笔交易卖家商户号
	ItemId           string `json:"item_id"`            //
}

// OrderStatus 订单状态 PROCESSING-处理中|SUCCESS-成功|FAIL-失败|TIMEOUT-超时
type OrderStatus string

const (
	PROCESSING OrderStatus = "PROCESSING"
	SUCCESS    OrderStatus = "SUCCESS"
	FAIL       OrderStatus = "FAIL"
	TIMEOUT    OrderStatus = "TIMEOUT"
)
