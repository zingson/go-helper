package douyinpoi

import (
	"github.com/sirupsen/logrus"
)

// GoodsDraftGet 查询商品草稿数据
// 接口文档地址:https://developer.open-douyin.com/docs/resource/zh-CN/dop/develop/openapi/life-service-open-ability/micro-app/goods-repo/goods.draft.get/
func GoodsDraftGet(token string) (res string) {
	url := GOODS_DRAFT_GET_URL + "?access_token=" + token + "&product_ids=7181731804364146721"
	result := CurlGet(url, "application/json", token)
	logrus.Infof("douyin查询商品草稿数据，url:%s, response:%s", url, result)
	//if err := json.Unmarshal([]byte(result), &res); err != nil {
	//	fmt.Printf("json unmarshal err:%v：", err)
	//}

	return res
}

//type GoodsDraftGetResponse struct {
//	Data struct {
//		ProductAttrs []struct {
//			Desc       string `json:"desc"`
//			IsMulti    bool   `json:"is_multi"`
//			IsRequired bool   `json:"is_required"`
//			Key        string `json:"key"`
//			Name       string `json:"name"`
//			ValueDemo  string `json:"value_demo"`
//			ValueType  string `json:"value_type"`
//		} `json:"product_attrs"`
//		SkuAttrs []struct {
//			Desc       string `json:"desc"`
//			IsMulti    bool   `json:"is_multi"`
//			IsRequired bool   `json:"is_required"`
//			Key        string `json:"key"`
//			Name       string `json:"name"`
//			ValueDemo  string `json:"value_demo"`
//			ValueType  string `json:"value_type"`
//		} `json:"sku_attrs"`
//	} `json:"data"`
//	Base struct {
//		LogID       string `json:"log_id"`
//		GatewayCode int    `json:"gateway_code"`
//		GatewayMsg  string `json:"gateway_msg"`
//		BizCode     int    `json:"biz_code"`
//		BizMsg      string `json:"biz_msg"`
//	} `json:"base"`
//}
