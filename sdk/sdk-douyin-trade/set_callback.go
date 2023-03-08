package helper

import (
	"encoding/json"
	"fmt"
)

// SetCallback 抖音开放平台交易系统设置回调地址:抖音开放平台交易系统在下单、退款、分账等流程中需要与开发者系统同步信息，这个接口为抖音开放平台交易系统同步信息给开发者提供回调地址。
// 接口文档地址:https://developer.open-douyin.com/docs/resource/zh-CN/mini-app/develop/server/trading/callback-config/config-callback-address
func SetCallback(req *SetCallbackRequest) (res *PubRes, err error) {
	url := SET_CALL_BACK_URL
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

type SetCallbackRequest struct {
	Create_order_callback    string `json:"create_order_callback"`
	Refund_callback          string `json:"refund_callback"`
	Delivery_qrcode_redirect string `json:"delivery_qrcode_redirect"`
	Book_callback            string `json:"book_callback"`
}
