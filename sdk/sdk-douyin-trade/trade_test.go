package helper

import (
	"dmkt/src/config"
	"encoding/json"
	"fmt"
	douyinpoi "github.com/zingson/go-helper/sdk/sdk-douyin-poi"
	"testing"
)

// 担保支付 订单同步
func TestOrderPush(t *testing.T) {
	token, _ := GetAccessToken(nil)
	item := &Item{
		Item_code: "",
		Img:       "",
		Title:     "放心充",
		Price:     1,
	}
	itemList := make([]*Item, 0)
	itemList = append(itemList, item)
	detail := &OrderDetail{
		OrderId:    "15366157291876147204",
		CreateTime: 1655192724000,
		Status:     "已核销",
		Amount:     1,
		TotalPrice: 1,
		DetailUrl:  "",
		Item_list:  itemList,
	}
	detailJson, err := json.Marshal(detail)
	if err != nil {
		fmt.Printf("json marshal err:%v", err)
	}
	douyinConf := config.Get[*douyinpoi.ClientTokenReq](config.DOUYIN_WJSQ) //抖音小程序 万家闪券 配置
	OrderPush(&OrderPushRequest{
		ClientKey:   douyinConf.ClientKey,
		AccessToken: token,
		AppName:     "douyin",
		OpenId:      "ee4aac0c-0ee9-4f93-b76d-e1298a3ac679",
		UpdateTime:  1655192724589,
		OrderType:   0,
		OrderStatus: 4,
		OrderDetail: string(detailJson),
	})
}

// 担保支付 结算分账
func TestOrderSettle(t *testing.T) {
	douyinConf := config.Get[*douyinpoi.ClientTokenReq](config.DOUYIN_WJSQ) //抖音小程序 万家闪券 配置
	OrderSettle(&OrderSettleRequest{
		AppId:       douyinConf.AppId,
		OutSettleNo: "st15366157291876147204",
		OutOrderNo:  "15366157291876147204",
		SettleDesc:  "分账",
	})
}

// 交易系统-回调设置-设置回调地址
func TestSetCallback(t *testing.T) {
	SetCallback(&SetCallbackRequest{
		Create_order_callback: "https://api.d.himkt.cn/gw/dmkt/v2/douyin/create_order_callback",
		Refund_callback:       "https://api.d.himkt.cn/gw/dmkt/v2/douyin/refund_callback",
	})
}

// 交易系统-回调设置-查询回调地址
func TestQuerySettings(t *testing.T) {
	QuerySettings()
}

// 交易系统-发起分账
func TestCreateSettle(t *testing.T) {
	CreateSettle(&CreateSettleRequest{
		OutOrderNo:  "15366157291876147204",
		OutSettleNo: "st15366157291876147204",
		SettleDesc:  "分账",
	})
}
