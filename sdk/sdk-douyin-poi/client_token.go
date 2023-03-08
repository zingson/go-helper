package douyinpoi

import (
	"encoding/json"
	"fmt"
)

// ClientToken 该接口用于获取接口调用的凭证client_access_token，主要用于调用不需要用户授权就可以调用的接口；该接口适用于抖音/头条授权。
// 生成client_token
// 接口文档地址:https://developer.open-douyin.com/docs/resource/zh-CN/dop/develop/openapi/account-permission/client-token
func ClientToken(req *ClientTokenRequest) (res *ClientTokenResponse) {
	url := REQUEST_ADDRESS
	reqMap := make(map[string]string)
	reqJson, err := json.Marshal(req)
	if err != nil {
		fmt.Printf("json marshal err:%v", err)
		return
	}
	err = json.Unmarshal(reqJson, &reqMap)
	if err != nil {
		fmt.Printf("json Unmarshal err:%v", err)
		return
	}
	result := CurlPostsForm(url, reqMap, "multipart/form-data", "", 2)
	if err = json.Unmarshal([]byte(result), &res); err != nil {
		fmt.Printf("json unmarshal err:%v：", err)
	}
	return
}

type ClientTokenReq struct {
	AppId        string `json:"appId"`        // 应用appId
	AppSecret    string `json:"appSecret"`    // 应用appSecret
	ClientKey    string `json:"clientKey"`    // 应用唯一标识
	ClientSecret string `json:"clientSecret"` // 应用唯一标识对应的密钥
	PrivateKey   string `json:"privateKey"`   // 应用私钥
	PublicKey    string `json:"publicKey"`    // 应用私钥
}
type ClientTokenRequest struct {
	ClientKey    string `json:"client_key"`    // 应用唯一标识
	ClientSecret string `json:"client_secret"` // 应用唯一标识对应的密钥
	GrantType    string `json:"grant_type"`    // 传client_credential
}

type ClientTokenResponse struct {
	Data    ClientTokenResponseData `json:"data"`
	Message string                  `json:"message"`
}
type ClientTokenResponseData struct {
	ExpiresIn   int64     `json:"expires_in"`   //access_token接口调用凭证超时时间，单位（秒)
	AccessToken string    `json:"access_token"` // 接口调用凭证
	Description string    `json:"description"`  // 错误码描述
	ErrorCode   errorCode `json:"error_code"`   // 错误码
}
