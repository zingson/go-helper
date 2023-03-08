package douyinpoi

import (
	"encoding/json"
	"fmt"
)

// SupplierQuerySupplier 店铺匹配状态查询
// 接口文档地址:https://open.douyin.com/platform/doc?doc=docs/openapi/life-service-open-ability/shop/status-query
func SupplierQuerySupplier(req *SupplierQuerySupplierRequest, token string) (res *SupplierQuerySupplierResponse) {
	url := SUPPLIER_QUERY_SUPPLIER_REQUEST_ADDRESS
	reqJson, err := json.Marshal(req)
	if err != nil {
		fmt.Printf("json marshal err:%v", err)
	}
	result := CurlPosts(url, string(reqJson), "application/json", token, 1)
	if err = json.Unmarshal([]byte(result), &res); err != nil {
		fmt.Printf("json unmarshal err:%v：", err)
	}
	return res
}

type SupplierQuerySupplierRequest struct {
	SupplierExtId string `json:"supplier_ext_id"` // 第三方店铺id列表，多个商铺id用 , 分割，单次查询最多50个店铺。
}
type SupplierQuerySupplierResponse struct {
	Data SupplierQuerySupplierResponseData `json:"data"`
}

type SupplierQuerySupplierResponseData struct {
	ErrorCode       int64                  `json:"error_code"`        // 错误码
	MatchResultList []*MatchResultListData `json:"match_result_list"` // 匹配的结果
	Description     string                 `json:"description"`       // 错误码描述
}

type MatchResultListData struct {
	Status        errorCode `json:"status"`          // 匹配状态，0-没有匹配，1-匹配中，2-匹配完成，3-匹配失败
	SupplierExtId string    `json:"supplier_ext_id"` // 商户ID
	TaskId        string    `json:"task_id"`         // 匹配任务ID
	PoiId         string    `json:"poi_id"`          // 抖音POIID

}
