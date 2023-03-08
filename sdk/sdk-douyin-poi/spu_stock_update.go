package douyinpoi

import (
	"encoding/json"
	"fmt"
)

//  SpuStockUpdate 多门店SPU库存同步(商品库存同步)
// 接口文档地址: https://open.douyin.com/platform/doc?doc=docs/openapi/life-service-open-ability/goods-repo/spu-repo-sync
func SpuStockUpdate(req *SpuStockUpdateRequest, token string) (res *SpuStockUpdateResponse) {
	url := SPU_STOCK_UPDATE_REQUEST_ADDRESS
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

type SpuStockUpdateRequest struct {
	SpuExtId string `json:"spu_ext_id"` //接入方商品ID
	Stock    int64  `json:"stock"`      //库存
}

type SpuStockUpdateResponse struct {
	Data  SpuStockUpdateResponseData `json:"data"`
	Extra interface{}                `json:"extra"`
}

type SpuStockUpdateResponseData struct {
	Description string    `json:"description"` // 	错误码描述
	ErrorCode   errorCode `json:"error_code"`  // 错误码
	SpuExtId    string    `json:"spu_ext_id"`  //	接入方SPU ID
}
