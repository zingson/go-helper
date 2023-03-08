package douyinpoi

import (
	"encoding/json"
	"fmt"
)

// SupplierQueryTask 店铺匹配任务结果查询
// 接口文档地址: https://open.douyin.com/platform/doc?doc=docs/openapi/life-service-open-ability/shop/task-query
func SupplierQueryTask(req *SupplierQueryTaskRequest, token string) (res *SupplierQueryTaskResponse) {
	url := SUPPLIER_QUERY_TASK_REQUEST_ADDRESS
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

type SupplierQueryTaskRequest struct {
	SupplierTaskIds string `json:"supplier_task_ids"` // 第三方上传任务id列表，多个任务id用 , 分割，单次查询最多10个任务。
}

type SupplierQueryTaskResponse struct {
	Description     string                                      `json:"description"`       // 错误码描述
	ErrorCode       errorCode                                   `json:"error_code"`        // 错误码
	MatchResultList []*SupplierQueryTaskResponseMatchResultList `json:"match_result_list"` // 匹配状态信息
}
type SupplierQueryTaskResponseMatchResultList struct {
	MismatchStatusDesc string `json:"mismatch_status_desc"` // 匹配状态描述
	Province           string `json:"province"`             // POI所在省份
	Address            string `json:"address"`              // POI地址
	MatchStatus        int64  `json:"match_status"`         // 匹配状态，0-等待匹配，1-正在匹配，2-匹配成功，3-匹配失败
	PoiId              string `json:"poi_id"`               // 高德POI ID/抖音POIID
	PoiName            string `json:"poi_name"`             // POI名称
	SupplierExtId      string `json:"supplier_ext_id"`      // 第三方商户ID
	City               string `json:"city"`                 // POI所在城市
	Extra              string `json:"extra"`                // 其他信息
}
