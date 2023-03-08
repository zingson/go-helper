package helper

import (
	"encoding/json"
	"fmt"
)

// CreateSettle 发起分账
// 接口文档地址 https://developer.open-douyin.com/docs/resource/zh-CN/mini-app/develop/server/trading/settlement/create-settle
func CreateSettle(req *CreateSettleRequest) (res *CreateSettleResponse, err error) {
	url := CREATE_SETTLE_URL
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

type CreateSettleRequest struct {
	OutOrderNo   string `json:"out_order_no"`            //开发者侧订单 id
	OutSettleNo  string `json:"out_settle_no"`           //开发者侧分账单 id
	SettleDesc   string `json:"settle_desc"`             //分账描述
	SettleParams string `json:"settle_params,omitempty"` // 其他分账方（除卖家之外的），长度 <= 512 字节 [{\"merchant_uid\":\"分账方商户号1\",\"amount\":100}]
	CpExtra      string `json:"cp_extra,omitempty"`      //开发者自定义透传字段
	NotifyUrl    string `json:"notify_url,omitempty"`    //分账结果通知地址，若不填，默认使用在https://developer.open-douyin.com/microapp/ttf4d2826f6becc24001/pay页面设置的支付回调地址。 ttf4d2826f6becc24001为小程序的 AppID。
}

type CreateSettleResponse struct {
	PubRes
	RespExtra struct {
		Logid string `json:"logid"`
	} `json:"resp_extra"`
	Data struct {
		SettleID string `json:"settle_id"`
	} `json:"data"`
}
