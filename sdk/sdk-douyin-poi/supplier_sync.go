package douyinpoi

import (
	"dmkt/src/biz/shp/consts"
	"encoding/json"
	"fmt"
	"github.com/sirupsen/logrus"
)

// SupplierSync 商铺同步
// 接口文档地址:https://open.douyin.com/platform/doc?doc=docs/openapi/life-service-open-ability/shop/synchronism
func SupplierSync(req *SupplierSyncRequest, token string) (res *SupplierSyncResponse) {
	url := SUPPLIER_SYNC_REQUEST_ADDRESS + "?access_token=" + token
	reqJson, err := json.Marshal(req)
	if err != nil {
		fmt.Printf("json marshal err:%v", err)
	}
	result := CurlPosts(url, string(reqJson), "application/json", token, 2)
	logrus.Infof("POI 商铺同步匹配，request:%s, response:%s", string(reqJson), result)
	if err = json.Unmarshal([]byte(result), &res); err != nil {
		fmt.Printf("json unmarshal err:%v：", err)
	}
	return res
}

type SupplierSyncRequest struct {
	//ContactPhone string    `json:"contact_phone"` // 联系手机号
	//ContactTel   string    `json:"contact_tel"`   // 联系座机号
	//Images       []*string `json:"images"`        // 店铺图片
	//MerchantUid string `json:"merchant_uid"` // 商户号；商家担保交易中的收款账户ID
	//ServiceProvider SyncServiceProviderData `json:"service_provider"` // 服务商资质信息
	SupplierExtId string `json:"supplier_ext_id"` // 接入方店铺id
	//Tags          []*string `json:"tags"`            // 标签
	//Latitude      string    `json:"latitude"`        // 纬度
	Status consts.PoiSupplierStatus `json:"status"` // 在线状态 1 - 在线; 2 - 下线
	//TypeCode      string    `json:"type_code"`       // POI品类编码
	//TypeName      string    `json:"type_name"`       // POI品类描述 eg. 美食;中式餐饮;小龙虾
	//Address       string    `json:"address"`         // 店铺地址
	//AvgCost       int64     `json:"avg_cost"`        // 人均消费（单位分）
	//CustomerInfo  SyncCustomerInfoData  `json:"customer_info"`   // 商家资质信息
	//Description string                `json:"description"` // 店铺介绍(<=500字)
	Name string `json:"name"` // 店铺名称
	//Recommends  []*SyncRecommendsData `json:"recommends"`  // 推荐
	//Services   []interface{} `json:"services"`   // 店铺提供的服务列表
	Attributes interface{} `json:"attributes"` // 店铺属性字段，编号规则：垂直行业 1xxx-酒店民宿 2xxx-餐饮 3xxx-景区 通用属性-9yxxx
	//Longitude   string                `json:"longitude"`   // 经度
	//OpenTime    []*string             `json:"open_time"`   // 营业时间, 从周一到周日，list长度为7，不营业则为空字符串
	PoiId string                 `json:"poi_id"` // 抖音poi id,三方如果使用高德poi id可以通过/poi/query/接口转换，其它三方poi id走poi匹配功能进行抖音poi id获取
	Type  consts.PoiSupplierType `json:"type"`   // 店铺类型 1 - 酒店民宿 2 - 餐饮 3 - 景区 4 - 电商 5 - 教育 6 - 丽人 7 - 爱车 8 - 亲子 9 - 宠物 10 - 家装 11 - 娱乐场所 12 - 图文快印
}

type SyncServiceProviderData struct {
	BusinessLicenseExtId string    `json:"business_license_ext_id"` // 服务商营业执照
	IndustryLicenseExtId []*string `json:"industry_license_ext_id"` // 服务商行业许可证
}
type SyncCustomerInfoData struct {
	PowerOfAttorney SyncPowerOfAttorneyData    `json:"power_of_attorney"` // 服务商和商家合作协议/授意书
	BusinessLicense SyncBusinessLicenseData    `json:"business_license"`  // 商家营业执照
	IndustryLicense []*SyncIndustryLicenseData `json:"industry_license"`  // 行业许可证
	OtherInfo       []*SyncOtherInfoData       `json:"other_info"`        // 其他补充材料
}
type SyncPowerOfAttorneyData struct {
	ExtId string `json:"ext_id"` // 合作协议/授意书外部id
	Url   string `json:"url"`    // 合作协议/授意书链接
}
type SyncBusinessLicenseData struct {
	Company string `json:"company"` //服务上营业执照公司名称
	ExtId   string `json:"ext_id"`  // 商家营业执照外部id
	Url     string `json:"url"`     // 商家营业执照链接
}
type SyncIndustryLicenseData struct {
	ExtId string `json:"ext_id"` // 商家营业执行业许可证外部id
	Url   string `json:"url"`    // 商家行业许可证链接
}
type SyncOtherInfoData struct {
	ExtId string `json:"ext_id"` // 其他补充材料外部id
	Url   string `json:"url"`    // 其他补充外部链接
}
type SyncRecommendsData struct {
	ImageUrl string `json:"image_url"` // 推荐内容链接(图片，或者视频链接）
	Title    string `json:"title"`     // 推荐描述
}
type SupplierSyncResponse struct {
	Data  SupplierSyncResponseData `json:"data"`
	Extra interface{}              `json:"extra"`
}
type SupplierSyncResponseData struct {
	Description string    `json:"description"` // 错误码描述
	ErrorCode   errorCode `json:"error_code"`  // 错误码
	SupplierId  string    `json:"supplier_id"` // 抖音平台商户ID
}
