package helper

import (
	"encoding/json"
	"fmt"
)

// QuerySettings 查询开发者设置的回调地址
// 接口文档地址:https://developer.open-douyin.com/docs/resource/zh-CN/mini-app/develop/server/trading/callback-config/query-callback-address
func QuerySettings() (res *PubRes, err error) {
	url := QUERY_SETTINGS_URL
	result, err := Request(nil, url, "POST", "")
	if err = json.Unmarshal([]byte(result), &res); err != nil {
		fmt.Printf("json unmarshal err:%v：", err)
	}
	return
}
