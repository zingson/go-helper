package douyinpoi

import (
	"encoding/json"
	"fmt"
	"github.com/sirupsen/logrus"
)

// SupplierQuery 查询店铺
// 接口文档地址: https://open.douyin.com/platform/doc?doc=docs/openapi/life-service-open-ability/shop/search
func SupplierQuery(supplier_ext_id string, token string) (res *SupplierQueryResponse) {
	//url := SUPPLIER_QUERY_REQUEST_ADDRESS
	//reqJson, err := json.Marshal(req)
	//if err != nil {
	//	fmt.Printf("json marshal err:%v", err)
	//}
	//result := CurlPosts(url, string(reqJson), "application/json", token, 1)
	//if err = json.Unmarshal([]byte(result), &res); err != nil {
	//	fmt.Printf("json unmarshal err:%v：", err)
	//}
	url := SUPPLIER_QUERY_REQUEST_ADDRESS + "?access_token=" + token + "&supplier_ext_id=" + supplier_ext_id
	result := CurlGet(url, "application/json", token)
	logrus.Infof("查询店铺，url:%s, response:%s", url, result)
	if err := json.Unmarshal([]byte(result), &res); err != nil {
		fmt.Printf("json unmarshal err:%v：", err)
	}
	return
}

type SupplierQueryRequest struct {
	SupplierExtId string `json:"supplier_ext_id"`
}
type SupplierQueryResponse struct {
	Data SupplierQueryResponseData `json:"data"`
}

type SupplierQueryResponseData struct {
	DataDetail  interface{}                 `json:"data_detail"`
	Description string                      `json:"description"` // 错误码描述
	ErrorCode   errorCode                   `json:"error_code"`  // 错误码
	SyncStatus  SupplierQuerySyncStatusData `json:"sync_status"` // 数据同步结果
}
type SupplierQuerySyncStatusData struct {
	LastSyncStatus int64  `json:"last_sync_status"` // 最近一次同步状态, 0 - 未处理; 1 - 通过; 2 - 未通过
	FailReason     string `json:"fail_reason"`      // 同步失败原因，抖音风控政策问题，该字段无法提供太多信息，目前审核不通过联系抖音运营做进一步处理
}
