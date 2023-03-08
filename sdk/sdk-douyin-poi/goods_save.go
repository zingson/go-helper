package douyinpoi

import (
	"encoding/json"
	"fmt"
	"github.com/sirupsen/logrus"
)

// GoodsSave 创建/更新商品
// 接口文档地址:https://bytedance.feishu.cn/docx/doxcncpgP3vi7QtK4CPX7Ark1sd
func GoodsSave(req *GoodsSaveRequest, token string) (res *GoodsSaveResponse) {
	url := GOODS_SAVE_URL + "?access_token=" + token
	reqJson, err := json.Marshal(req)
	if err != nil {
		fmt.Printf("json marshal err:%v", err)
	}
	result := CurlPosts(url, string(reqJson), "application/json", token, 2)
	logrus.Infof("-----创建/更新商品，url:%s, request:%s, response:%s", url, string(reqJson), result)
	if err = json.Unmarshal([]byte(result), &res); err != nil {
		fmt.Printf("json unmarshal err:%v：", err)
	}

	return res
}

type GoodsSaveResponse struct {
	Data struct {
		ProductId string `json:"product_id"`
	} `json:"data"`
	Base struct {
		LogID       string `json:"log_id"`
		GatewayCode int    `json:"gateway_code"`
		GatewayMsg  string `json:"gateway_msg"`
		BizCode     int    `json:"biz_code"`
		BizMsg      string `json:"biz_msg"`
	} `json:"base"`
}

type GoodsSaveRequest struct {
	Product *Product `json:"product"`
	Sku     *Sku     `json:"sku"`
	//OwnerAccountId string   `json:"owner_account_id"` //商品归属账户ID，非必传；传入时须与该商家满足商服关系；没有还没有商户ID，可以先不传
}

type Product struct {
	ProductId        string              `json:"product_id,omitempty"`         //商品Id，创建时不必填写，更新时如有out_id可不填写
	OutID            string              `json:"out_id"`                       //外部商品id
	ProductName      string              `json:"product_name"`                 //商品名
	CategoryFullName string              `json:"category_full_name,omitempty"` //品类全名，保存时不必填写
	CategoryID       int64               `json:"category_id"`                  //品类id
	ProductType      int                 `json:"product_type"`                 //商品类型：1 : 团购套餐 3 : 预售券 4 : 日历房 5 : 门票 7 : 旅行跟拍 8 : 一日游 11 : 代金券 14: x项目
	BizLine          int                 `json:"biz_line"`                     //业务线 1-闭环自研开发者 3-直连服务商 5-小程序
	AccountName      string              `json:"account_name"`                 //商家名
	SoldStartTime    int64               `json:"sold_start_time"`              //售卖开始时间
	SoldEndTime      int64               `json:"sold_end_time"`                //售卖结束时间
	OutUrl           string              `json:"out_url,omitempty"`            //第三方跳转链接，小程序商品必填
	PoiList          []*Poi              `json:"poi_list"`                     //店铺列表
	AttrKeyValueMap  *ProAttrKeyValueMap `json:"attr_key_value_map,omitempty"` //商品属性KV
	//Telephone []string `json:"telephone,omitempty"`
}

type Sku struct {
	SkuId           string              `json:"sku_id,omitempty"`
	SkuName         string              `json:"sku_name"`
	OriginAmount    int                 `json:"origin_amount,omitempty"` //原价
	ActualAmount    int                 `json:"actual_amount"`           //实际支付价格
	Stock           *Stock              `json:"stock"`
	OutSkuId        string              `json:"out_sku_id"` //第三方id
	Status          int                 `json:"status"`     //状态 1-在线，2-待售 ； 默认传1
	AttrKeyValueMap *SkuAttrKeyValueMap `json:"attr_key_value_map,omitempty"`
}

type Poi struct {
	SupplierExtID string `json:"supplier_ext_id,omitempty"` //接入方店铺id，保存时必传
}

type Stock struct {
	LimitType int   `json:"limit_type"` //库存上限类型，为2时stock_qty和avail_qty字段无意义  1-有限库存 2-无限库存
	StockQty  int64 `json:"stock_qty"`  //总库存，limit_type=2时无意义
}

type SkuAttrKeyValueMap struct {
	CodeSourceType string `json:"code_source_type"`
	//Commodity                 string `json:"commodity,omitempty"`
	LimitRule                 string `json:"limit_rule"`
	RefundNeedMerchantConfirm string `json:"refund_need_merchant_confirm"`
	SettleType                string `json:"settle_type"`
	UseType                   string `json:"use_type"`
}

type ProAttrKeyValueMap struct {
	Appointment  string `json:"appointment"`
	AutoRenew    string `json:"auto_renew"`
	CanNoUseDate string `json:"can_no_use_date"`
	//EnvironmentImageList      string `json:"environment_image_list,omitempty"`
	ImageList                 string `json:"image_list"`
	ParkingLotWithin1Km       string `json:"parking_lot_within_1km"`
	RealNameInfo              string `json:"real_name_info"`
	RecPersonNum              string `json:"rec_person_num"`
	ShowChannel               string `json:"show_channel"`
	SuperimposedDiscounts     string `json:"superimposed_discounts"`
	UseDate                   string `json:"use_date"`
	UseTime                   string `json:"use_time"`
	BringOutMeal              string `json:"bring_out_meal"`
	FreePack                  string `json:"free_pack"`
	Notification              string `json:"Notification"`
	PrivateRoom               string `json:"private_room"`
	RecPersonNumMax           string `json:"rec_person_num_max"`
	RefundPolicy              string `json:"RefundPolicy"`
	RefundNeedMerchantConfirm string `json:"refund_need_merchant_confirm"`
	Description               string `json:"Description"`
	LimitUseRule              string `json:"limit_use_rule"`
	EntryType                 string `json:"EntryType"`
}

type EntrySchema struct {
	AppId  string `json:"appId"`            //订单详情页路径，没有前导的/，该字段不能为空，长度 <= 512byte pages/xxxindexxxx
	Path   string `json:"path"`             //订单详情页路径，没有前导的/，该字段不能为空，长度 <= 512byte pages/xxxindexxxx
	Params string `json:"params,omitempty"` //路径参数，自定义的 json 结构，序列化成字符串存入该字段，平台不限制，但是写入的内容需要能够保证生成访问订单详情的 schema 能正确跳转到小程序内部的订单详情页，长度 <= 512byte  {\"id\":\"xxxxxx\"}
}
