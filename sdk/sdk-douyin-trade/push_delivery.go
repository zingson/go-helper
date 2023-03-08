package helper

import (
	"encoding/json"
	"fmt"
)

// PushDelivery 同步订单的核销状态到开放平台。
// 接口文档地址https://developer.open-douyin.com/docs/resource/zh-CN/mini-app/develop/server/trading/write-off/third-party-code/push-delivery
func PushDelivery(req *PushDeliveryRequest) (res *PubRes, err error) {
	url := PUSH_DELIVERY_URL
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

type PushDeliveryRequest struct {
	OutOrderNo    string `json:"out_order_no"`
	ItemOrderList []struct {
		ItemOrderID string `json:"item_order_id"`
	} `json:"item_order_list,omitempty"`
	UseAll  bool   `json:"use_all,omitempty"`
	PoiInfo string `json:"poi_info,omitempty"`
}
