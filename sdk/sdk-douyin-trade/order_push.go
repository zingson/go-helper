package helper

import (
	"encoding/json"
	"fmt"
	douyinpoi "github.com/zingson/go-helper/sdk/sdk-douyin-poi"
)

// OrderPush 担保交易的订单同步（如果使用担保交易的支付，需要调用此接口同步订单）。
// 接口文档地址 https://developer.open-douyin.com/docs/resource/zh-CN/mini-app/develop/server/ecpay/order/order-sync/
func OrderPush(conf *douyinpoi.ClientTokenReq, req *OrderPushRequest) (res *PubRes, err error) {
	url := ORDER_SYNC_URL
	reqJson, err := json.Marshal(req)
	if err != nil {
		fmt.Printf("json marshal err:%v", err)
	}
	result, err := Request(conf, url, "POST", string(reqJson))
	if err = json.Unmarshal([]byte(result), &res); err != nil {
		return
	}
	return
}

type OrderPushRequest struct {
	ClientKey   string `json:"client_key,omitempty"`
	AccessToken string `json:"access_token"`
	ExtShopId   string `json:"ext_shop_id,omitempty"`
	AppName     string `json:"app_name"`
	OpenId      string `json:"open_id"`
	UpdateTime  int64  `json:"update_time"`
	OrderDetail string `json:"order_detail"`
	OrderType   int64  `json:"order_type"`
	OrderStatus int64  `json:"order_status"`
	Extra       string `json:"extra,omitempty"`
}

type OrderDetail struct {
	OrderId    string  `json:"order_id"`
	CreateTime int64   `json:"create_time"`
	Status     string  `json:"status"`
	Amount     int64   `json:"amount"`
	TotalPrice int64   `json:"total_price"`
	DetailUrl  string  `json:"detail_url"`
	Item_list  []*Item `json:"item_list"`
}

type Item struct {
	Item_code string `json:"item_code"`
	Img       string `json:"img"`
	Title     string `json:"title"`
	Sub_title string `json:"sub_title,omitempty"`
	Amount    int64  `json:"amount,omitempty"`
	Price     int64  `json:"price"`
}
