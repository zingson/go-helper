package helper

import (
	"encoding/json"
	"fmt"
)

// CreateOrder 预下单
// 接口文档地址 https://developer.open-douyin.com/docs/resource/zh-CN/mini-app/develop/server/trading/pre-order/create-order
func CreateOrder(req *CreateOrderRequest) (res *CreateOrderResponse, err error) {
	url := CREATE_ORDER_URL
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

type CreateOrderRequest struct {
	GoodsList   []*Goods `json:"goods_list"`   //商品信息 //POI 商品会从商品库里查询商品信息，不会使用开发者传的数据。
	TotalAmount int      `json:"total_amount"` //订单总价，单位分
	//DiscountAmount int    `json:"discount_amount,omitempty"` //折扣金额，单位分
	AppID string `json:"app_id"`
	//PhoneNum         string `json:"phone_num,omitempty"` //用户手机号
	//ContactName      string `json:"contact_name,omitempty"`
	Extra        string `json:"extra,omitempty"`          //下单备注信息
	OpenID       string `json:"open_id"`                  //用户 OpenID
	OutOrderNo   string `json:"out_order_no"`             //开发者的单号
	PayNotifyURL string `json:"pay_notify_url,omitempty"` //支付结果通知地址
	//PayExpireSeconds int    `json:"pay_expire_seconds,omitempty"`//支付超时时间，单位秒，例如 300 表示 300 秒后过期；不传或传 0 会使用默认值 300。
	OrderEntrySchema *OrderEntrySchema `json:"order_entry_schema"` //订单详情页信息
	CpExtra          string            `json:"cp_extra,omitempty"` //开发者自定义透传字段
	//PriceCalculationDetail struct {
	//	CalculationType     int `json:"calculation_type"`
	//	OrderDiscountDetail struct {
	//		GoodsTotalDiscountAmount int `json:"goods_total_discount_amount"`
	//		MarketingDetailInfo      []struct {
	//			//DiscountAmount int    `json:"discount_amount,omitempty"` //折扣金额，单位分
	//			DiscountRange int    `json:"discount_range"`
	//			ID            string `json:"id"`
	//			Note          string `json:"note"`
	//			Subtype       string `json:"subtype"`
	//			Title         string `json:"title"`
	//			Type          int    `json:"type"`
	//		} `json:"marketing_detail_info"`
	//		OrderTotalDiscountAmount int `json:"order_total_discount_amount"`
	//	} `json:"order_discount_detail"`
	//	GoodsDiscountDetail []struct {
	//		GoodsID             string `json:"goods_id"`
	//		MarketingDetailInfo []struct {
	//			DiscountAmount int    `json:"discount_amount"`
	//			DiscountRange  int    `json:"discount_range"`
	//			ID             string `json:"id"`
	//			Note           string `json:"note"`
	//			Subtype        string `json:"subtype"`
	//			Title          string `json:"title"`
	//			Type           int    `json:"type"`
	//		} `json:"marketing_detail_info"`
	//		Quantity            int `json:"quantity"`
	//		TotalAmount         int `json:"total_amount"`
	//		TotalDiscountAmount int `json:"total_discount_amount"`
	//	} `json:"goods_discount_detail"`
	//	ItemDiscountDetail []struct {
	//		GoodsID             string `json:"goods_id"`
	//		MarketingDetailInfo []struct {
	//			DiscountAmount int    `json:"discount_amount"`
	//			DiscountRange  int    `json:"discount_range"`
	//			ID             string `json:"id"`
	//			Note           string `json:"note"`
	//			Subtype        string `json:"subtype"`
	//			Title          string `json:"title"`
	//			Type           int    `json:"type"`
	//		} `json:"marketing_detail_info"`
	//		TotalAmount         int `json:"total_amount"`
	//		TotalDiscountAmount int `json:"total_discount_amount"`
	//	} `json:"item_discount_detail"`
	//} `json:"price_calculation_detail"` //营销算价结果信息
}

type Goods struct {
	Quantity    int    `json:"quantity"`              //商品数量
	Price       int    `json:"price,omitempty"`       //商品价格，单位（分） 注意：非 POI 商品必传
	GoodsTitle  string `json:"goods_title,omitempty"` //商品标题/商品名称
	GoodsImage  string `json:"goods_image,omitempty"` //商品图片链接
	Labels      string `json:"labels,omitempty"`      //商品标签，最多设置三个标签，例如：随时退｜免预约｜提前3日预约（“｜”是中文类型）
	GoodsID     string `json:"goods_id"`              //商品 id
	GoodsIDType int    `json:"goods_id_type"`         //商品 id 类别， POI 商品传 1；非 POI 商品传 2
	//DiscountAmount int    `json:"discount_amount,omitempty"` //折扣金额，单位分
	//DateRule string `json:"date_rule,omitempty"` //使用规则，如 “周一至周日可用”、“周一至周五可用”、“非节假日可用”，默认“周一至周日可用”
	//GoodsPage struct {
	//	Path   string `json:"path"`   //商品详情页路径
	//	Params string `json:"params"` //商品详情页路径参数
	//} `json:"goods_page,omitempty"` //商品详情页
	//OrderValidTime struct {
	//	ValidDuration int `json:"valid_duration"`
	//} `json:"order_valid_time,omitempty"` //券的有效期，详情见 order_valid_time 字段 注意： 非 POI 商品必传，POI 商品会从 POI 库里查询有效期信息，不会使用开发者传的数据。 如果是非 POI 商品，每个 goods_id 都要传券的有效期信息，否则会下单失败。
	//GoodsBookInfo struct {
	//	BookType     int `json:"book_type"`
	//	CancelPolicy int `json:"cancel_policy"`
	//} `json:"goods_book_info,omitempty"` //预约信息
	//MerchantUID string `json:"merchant_uid,omitempty"`//开发者自定义收款商户号，须申请白名单
}

type CreateOrderResponse struct {
	PubRes
	RespExtra struct {
		Logid string `json:"logid"`
	} `json:"resp_extra"`
	Data struct {
		OrderID           string     `json:"order_id"`        //抖音开平侧生成的订单号
		OutOrderNo        string     `json:"out_order_no"`    //开发者系统生成的订单号
		PayOrderID        string     `json:"pay_order_id"`    //吊起收银台的支付订单号
		PayOrderToken     string     `json:"pay_order_token"` //吊起收银台的 token
		ItemOrderInfoList []struct { //商品 item_order 信息
			ItemOrderIDList []string   `json:"item_order_id_list"` //item_order_id 列表，id 个数与下单时对应 goods_id 的 quantity 一致
			GoodsID         string     `json:"goods_id"`           //商品 id
			ItemOrderDetail []struct { //商品 item_order 详细信息
				ItemOrderID string `json:"item_order_id"` //item 单 id
				Price       int    `json:"price"`         //商品优惠后价格
			} `json:"item_order_detail"`
		} `json:"item_order_info_list"`
	} `json:"data"`
}
