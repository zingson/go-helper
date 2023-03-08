package helper

import (
	"encoding/json"
	"fmt"
)

// QueryOrder 查询订单信息
// 接口文档地址 https://developer.open-douyin.com/docs/resource/zh-CN/mini-app/develop/server/trading/pre-order/query-order
func QueryOrder(req *QueryOrderRequest) (res *QueryOrderResponse, err error) {
	url := QUERY_ORDER_URL
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

type QueryOrderRequest struct {
	OrderId    string `json:"order_id,omitempty"`     //抖音开平内部交易订单号，通过预下单回调传给开发者服务
	OutOrderNo string `json:"out_order_no,omitempty"` //开发者系统生成的订单号，与唯一order_id关联
}

type QueryOrderResponse struct {
	PubRes
	RespExtra struct {
		Logid string `json:"logid"`
	} `json:"resp_extra"`
	Data struct {
		OrderId        string `json:"order_id"`         //抖音开平侧订单号
		OutOrderNo     string `json:"out_order_no"`     //开发者侧订单号，与 order_id 一一对应
		RefundAmount   int64  `json:"refund_amount"`    //已退款金额，单位分
		Settle_amount  int64  `json:"settle_amount"`    //已分账金额，单位分
		TotalFee       int64  `json:"total_fee"`        //订单实际支付金额，单位[分]
		OrderStatus    string `json:"order_status"`     //订单状态， INIT： 初始状态 PROCESS： 订单处理中 SUCESS：成功 FAIL：失败 TIMEOUT：用户超时未支付
		PayTime        string `json:"pay_time"`         //支付时间，格式：2021-12-12 00:00:00
		PayChannel     int    `json:"pay_channel"`      //支付渠道枚举 1：微信， 2：支付宝 10：抖音支付
		ChannelPayId   string `json:"channel_pay_id"`   //渠道支付单号，如微信的支付单号
		SellerUid      string `json:"seller_uid"`       //卖家商户号 id
		ItemId         string `json:"item_id"`          //视频id
		CpExtra        string `json:"cp_extra"`         //预下单时开发者定义的透传信息
		Message        string `json:"message"`          //结果描述信息，如失败原因
		PaymentOrderId string `json:"payment_order_id"` //担保支付单 id
		DeliveryType   int    `json:"delivery_type"`    //订单核销类型： 0: 非闭环核销，开发者自行处理券码生成及展示，通过 push_delivery 接口推送核销状态。 1: 闭环核销，开平负责生券，开发者使用核销组件展示，使用验券准备和验券接口核销。
	} `json:"data"`
}
