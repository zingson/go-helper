package douyinpoi

import (
	"dmkt/src/biz/shp/consts"
	"encoding/json"
	"fmt"
)

// SpuStatusSync 多门店SPU状态同步(主要商品下线)
// 接口文档地址: https://open.douyin.com/platform/doc?doc=docs/openapi/life-service-open-ability/goods-repo/spu-status-sync
func SpuStatusSync(req *SpuStatusSyncRequest, token string) (res *SpuStatusSyncResponse) {
	url := SPU_STATUS_SYNC_REQUEST_ADDRESS
	reqJson, err := json.Marshal(req)
	if err != nil {
		fmt.Printf("json marshal err:%v", err)
	}
	result := CurlPosts(url, string(reqJson), "application/json", token, 2)
	if err = json.Unmarshal([]byte(result), &res); err != nil {
		fmt.Printf("json unmarshal err:%v：", err)
	}
	return
}

type SpuStatusSyncRequest struct {
	SpuExtIdList []string            `json:"spu_ext_id_list"` //	接入方商品ID列表
	Status       consts.PoiSpuStatus `json:"status"`          //	SPU状态， 1 - 在线; 2 - 下线
}

type SpuStatusSyncResponse struct {
	Data  SpuStatusSyncResponseData `json:"data"`
	Extra interface{}               `json:"extra"`
}

type SpuStatusSyncResponseData struct {
	Description  string    `json:"description"`     // 	错误码描述
	ErrorCode    errorCode `json:"error_code"`      // 错误码
	SpuExtIdList []*string `json:"spu_ext_id_list"` //	状态同步成功的spu_ext_id列表
}
