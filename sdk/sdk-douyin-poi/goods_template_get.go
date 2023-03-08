package douyinpoi

import (
	"encoding/json"
	"fmt"
	"github.com/sirupsen/logrus"
)

// GoodsTemplateGet 查询商品模板
// 接口文档地址:https://bytedance.feishu.cn/docx/doxcncpgP3vi7QtK4CPX7Ark1sd
func GoodsTemplateGet(token string) (res *GoodsTemplateGetResponse) {
	url := GOODS_TEMPLATE_GET_URL + "?access_token=" + token + "&category_id=19005002&product_type=3"
	result := CurlGet(url, "application/json", token)
	logrus.Infof("查询商品模板，url:%s, response:%s", url, result)
	if err := json.Unmarshal([]byte(result), &res); err != nil {
		fmt.Printf("json unmarshal err:%v：", err)
	}

	return res
}

type GoodsTemplateGetResponse struct {
	Data struct {
		ProductAttrs []struct {
			Desc       string `json:"desc"`
			IsMulti    bool   `json:"is_multi"`
			IsRequired bool   `json:"is_required"`
			Key        string `json:"key"`
			Name       string `json:"name"`
			ValueDemo  string `json:"value_demo"`
			ValueType  string `json:"value_type"`
		} `json:"product_attrs"`
		SkuAttrs []struct {
			Desc       string `json:"desc"`
			IsMulti    bool   `json:"is_multi"`
			IsRequired bool   `json:"is_required"`
			Key        string `json:"key"`
			Name       string `json:"name"`
			ValueDemo  string `json:"value_demo"`
			ValueType  string `json:"value_type"`
		} `json:"sku_attrs"`
	} `json:"data"`
	Base struct {
		LogID       string `json:"log_id"`
		GatewayCode int    `json:"gateway_code"`
		GatewayMsg  string `json:"gateway_msg"`
		BizCode     int    `json:"biz_code"`
		BizMsg      string `json:"biz_msg"`
	} `json:"base"`
}
