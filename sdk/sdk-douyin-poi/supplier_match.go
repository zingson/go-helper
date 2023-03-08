package douyinpoi

import (
	"encoding/json"
	"fmt"
	"github.com/sirupsen/logrus"
)

type SupplierMatchIsSuccess int64

const (
	SUPPLIER_MATCH_IS_SUCCESS_1 SupplierMatchIsSuccess = 1 //成功
	SUPPLIER_MATCH_IS_SUCCESS_2 SupplierMatchIsSuccess = 2 //失败
)

// SupplierMatch 发起店铺匹配POI同步任务
// 接口文档地址:https://open.douyin.com/platform/doc?doc=docs/openapi/life-service-open-ability/shop/poi-sync-task
func SupplierMatch(req *SupplierMatchRequest, token string) (res *SupplierMatchResponse) {
	url := SUPPLIER_MATACH_REQUEST_ADDRESS + "?access_token=" + token
	reqJson, err := json.Marshal(req)
	if err != nil {
		fmt.Printf("json marshal err:%v", err)
	}
	result := CurlPosts(url, string(reqJson), "application/json", token, 2)
	logrus.Infof("POI匹配，request:%s, response:%s", string(reqJson), result)
	if err = json.Unmarshal([]byte(result), &res); err != nil {
		fmt.Printf("json unmarshal err:%v：", err)
	}

	return res
}

type SupplierMatchRequest struct {
	MatchDataList []*SupplierMatchRequestData `json:"match_data_list"`
}

type SupplierMatchRequestData struct {
	City          string  `json:"city"`            // POI所在城市
	Latitude      float64 `json:"latitude"`        // 纬度
	Longitude     float64 `json:"longitude"`       // 经度
	AmapId        string  `json:"amap_id"`         // 高德POI ID
	Extra         string  `json:"extra"`           // 其他信息
	PoiId         string  `json:"poi_id"`          // 高德POI ID/抖音POIID
	PoiName       string  `json:"poi_name"`        // POI名称
	Province      string  `json:"province"`        // POI所在省份
	SupplierExtId string  `json:"supplier_ext_id"` // 第三方商户ID
	Address       string  `json:"address"`         // POI地址
}

type SupplierMatchResponse struct {
	Data SupplierMatchResponseData `json:"data"`
}

type SupplierMatchResponseData struct {
	ErrorCode   errorCode              `json:"error_code"`  // 错误码
	IsSuccess   SupplierMatchIsSuccess `json:"is_success"`  // 上传状态(1:成功，2:失败)
	TaskId      string                 `json:"task_id"`     // 抖音平台任务ID
	Description string                 `json:"description"` // 错误码描述
}
