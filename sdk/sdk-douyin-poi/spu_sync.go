package douyinpoi

import (
	"dmkt/src/biz/shp/consts"
	"encoding/json"
	"fmt"
	"github.com/sirupsen/logrus"
	"time"
)

// SpuSync 多门店SPU同步(上线商品)
// 接口文档地址:https://open.douyin.com/platform/doc?doc=docs/openapi/life-service-open-ability/goods-repo/spu-sync
func SpuSync(req *SpuSyncRequest, token string) (res *SpuSyncResponse) {
	bmilli := time.Now().UnixMilli()
	url := SPU_SYNC_REQUEST_ADDRESS
	reqJson, err := json.Marshal(req)
	if err != nil {
		fmt.Printf("json marshal err:%v", err)
	}
	result := CurlPosts(url, string(reqJson), "application/json", token, 2)
	if err = json.Unmarshal([]byte(result), &res); err != nil {
		fmt.Printf("json unmarshal err:%v：", err)
	}
	millisecond := fmt.Sprintf("%d", time.Now().UnixMilli()-bmilli)
	logrus.WithField("millisecond", millisecond).Infof("sdk douyin 商品同步 请求URI：%s  请求报文：%s  响应报文：%s  | %sms", url, string(reqJson), result, millisecond)
	return
}

type SpuSyncRequest struct {
	EntryInfo SpuSyncRequestEntryInfo `json:"entry_info"` // 入口信息
	//FrontCategoryTag   []*string   `json:"front_category_tag"`    // 前台品类标签
	SupplierExtIdList []string `json:"supplier_ext_id_list"` // 接入方店铺ID列表
	Price             int64    `json:"price"`                // 价格，单位分
	SpuExtId          string   `json:"spu_ext_id"`           // 接入方商品ID
	ImageList         []string `json:"image_list"`           // SPU图片，预售券必传
	//MpSettleType       int64       `json:"mp_settle_type"`        // 小程序结算类型 1-包销 2-代销
	OriginPrice int64 `json:"origin_price"` // 价格，单位分
	//TakeRate           string      `json:"take_rate"`             // 商品的抽佣率，万分数
	Attribute SpuSyncResponseAttribute `json:"attribute"` // SPU属性字段
	//Highlights         interface{} `json:"highlights"`            // 商品亮点标签
	Name string `json:"name"` // 商品名
	//OrderDependsOnDate bool        `json:"order_depends_on_date"` // i单是否依赖日期
	SpuType consts.PoiSpuType   `json:"spu_type"` // spu类型号，1-日历房，30-酒店民宿预售券，90-门票，91-团购券
	Status  consts.PoiSpuStatus `json:"status"`   // SPU状态， 1 - 在线; 2 - 下线
	Stock   int64               `json:"stock"`    // 库存
}
type SpuSyncRequestEntryInfo struct {
	EntryMiniApp SpuSyncRequestEntryMiniApp `json:"entry_miniApp"` //小程序入口参数
	EntryType    int64                      `json:"entry_type"`    //入口类型(1:H5，2:抖音小程序，3:抖音链接)
	EntryUrl     string                     `json:"entry_url"`     //入口链接
}
type SpuSyncRequestEntryMiniApp struct {
	Params string `json:"params"` //服务参数json
	Path   string `json:"path"`   //服务路径
	AppId  string `json:"app_id"` //小程序的appid
}
type SpuSyncResponse struct {
	Data  SpuSyncResponseData `json:"data"`
	Extra interface{}         `json:"extra"`
}
type SpuSyncResponseData struct {
	Description string    `json:"description"` // 	错误码描述
	ErrorCode   errorCode `json:"error_code"`  // 错误码
	SpuId       string    `json:"spu_id"`      //	抖音平台SPU ID
}

type SpuSyncResponseAttribute struct {
	A9001 SpuSyncResponseAttribute9001 `json:"9001"`
	//A9101 SpuSyncResponseAttribute9101 `json:"9101"`
}
type SpuSyncResponseAttribute9001 struct {
	IsConfirmImme bool `json:"is_confirm_imme"`
	IsNeedPick    bool `json:"is_need_pick"`
}
type SpuSyncResponseAttribute9101 struct {
	GrouponValidStart string `json:"groupon_valid_start"`
	OrderLimit        int64  `json:"order_limit"`
	OrderValidStart   string `json:"order_valid_start"`
}
