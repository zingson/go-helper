package douyinpoi

import (
	"encoding/json"
	"fmt"
)

// SpuGet 多门店SPU信息查询
// 接口文档地址: https://open.douyin.com/platform/doc?doc=docs/openapi/life-service-open-ability/goods-repo/spu-info-query
func SpuGet(req *SpuGetRequest, token string) (res *SpuGetResponse) {
	url := SPU_GET_REQUEST_ADDRESS
	//reqJson, err := json.Marshal(req)
	//if err != nil {
	//	fmt.Printf("json marshal err:%v", err)
	//}
	result := CurlGet(url, "application/json", token)
	if err := json.Unmarshal([]byte(result), &res); err != nil {
		fmt.Printf("json unmarshal err:%v：", err)
	}
	return
}

type SpuGetRequest struct {
	SpuExtId                   string    `json:"spu_ext_id"`
	NeedSpuDraft               bool      `json:"need_spu_draft"`
	SpuDraftCount              int64     `json:"spu_draft_count"`
	SupplierIdsForFilterReason []*string `json:"supplier_ids_for_filter_reason"`
}

type SpuGetResponse struct {
	Data SpuGetResponseData `json:"data"`
}

type SpuGetResponseData struct {
	Description string         `json:"description"` // 错误码描述
	ErrorCode   errorCode      `json:"error_code"`  // 错误码
	SpuDetail   interface{}    `json:"spu_detail"`
	SpuDraft    []*interface{} `json:"spu_draft"`
}
