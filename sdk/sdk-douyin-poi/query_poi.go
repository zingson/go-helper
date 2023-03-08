package douyinpoi

import (
	"encoding/json"
	"fmt"
	"github.com/sirupsen/logrus"
)

// QueryPoi 获取抖音 POI ID
// 接口文档地址:https://developer.open-douyin.com/docs/resource/zh-CN/dop/develop/openapi/life-service-open-ability/micro-app/shop/get-douyin-poi-id
func QueryPoi(amap_id string, token string) (res *QueryPoiRes) {
	url := QUERY_POI_REQUEST_ADDRESS + "?access_token=" + token + "&amap_id=" + amap_id
	result := CurlGet(url, "application/json", token)
	logrus.Infof("douyin获取抖音 POI ID，url:%s, response:%s", url, result)
	if err := json.Unmarshal([]byte(result), &res); err != nil {
		fmt.Printf("json unmarshal err:%v：", err)
	}

	return res
}

type QueryPoiRes struct {
	Data struct {
		Address     string `json:"address"`
		AmapID      string `json:"amap_id"`
		City        string `json:"city"`
		Description string `json:"description"`
		ErrorCode   string `json:"error_code"`
		Latitude    string `json:"latitude"`
		Longitude   string `json:"longitude"`
		PoiID       string `json:"poi_id"`
		PoiName     string `json:"poi_name"`
	} `json:"data"`
}
