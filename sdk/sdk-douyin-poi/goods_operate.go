package douyinpoi

import (
	"encoding/json"
	"fmt"
	"github.com/sirupsen/logrus"
)

// GoodsOperate 上架/下架商品
// 接口文档地址:https://bytedance.feishu.cn/docx/doxcncpgP3vi7QtK4CPX7Ark1sd
func GoodsOperate(req *GoodsOperateRequest, token string) (res *GoodsOperateResponse) {
	url := GOODS_OPERATE_URL + "?access_token=" + token
	reqJson, err := json.Marshal(req)
	if err != nil {
		fmt.Printf("json marshal err:%v", err)
	}
	result := CurlPosts(url, string(reqJson), "application/json", token, 2)
	logrus.Infof("-----创建/更新商品，url:%s, request:%s, response:%s", url, string(reqJson), result)
	if err = json.Unmarshal([]byte(result), &res); err != nil {
		fmt.Printf("json unmarshal err:%v：", err)
	}

	return res
}

type GoodsOperateResponse struct {
	Data struct {
		ProductId string `json:"product_id"`
	} `json:"data"`
	Base struct {
		LogID       string `json:"log_id"`
		GatewayCode int    `json:"gateway_code"`
		GatewayMsg  string `json:"gateway_msg"`
		BizCode     int    `json:"biz_code"`
		BizMsg      string `json:"biz_msg"`
	} `json:"base"`
}

type GoodsOperateRequest struct {
	ProductId string `json:"product_id"` //抖音商品id
	OutId     string `json:"out_id"`     //第三方商品id
	OpType    int64  `json:"op_type"`    //1 上线；2下线
}
